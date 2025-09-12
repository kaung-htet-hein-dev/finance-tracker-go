package services

import (
	"finance-tracker-go/internal/database"
	"finance-tracker-go/internal/models"
	"fmt"
	"time"
)

type InsightsService struct{}

func NewInsightsService() *InsightsService {
	return &InsightsService{}
}

func (s *InsightsService) GetInsights(userID uint) (*models.InsightResponse, error) {
	db := database.GetDB()

	// Get total income and expenses
	var totalIncome, totalExpenses float64
	
	db.Model(&models.Transaction{}).
		Where("user_id = ? AND type = ?", userID, models.Income).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&totalIncome)

	db.Model(&models.Transaction{}).
		Where("user_id = ? AND type = ?", userID, models.Expense).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&totalExpenses)

	// Get top spending categories
	var categoryInsights []models.CategoryInsight
	db.Model(&models.Transaction{}).
		Select("category, SUM(amount) as amount, COUNT(*) as count").
		Where("user_id = ? AND type = ?", userID, models.Expense).
		Group("category").
		Order("amount DESC").
		Limit(5).
		Scan(&categoryInsights)

	// Get monthly trends (last 6 months)
	monthlyTrends := s.getMonthlyTrends(userID)

	// Calculate savings rate
	savingsRate := float64(0)
	if totalIncome > 0 {
		savingsRate = ((totalIncome - totalExpenses) / totalIncome) * 100
	}

	// Calculate average expense
	var expenseCount int64
	db.Model(&models.Transaction{}).
		Where("user_id = ? AND type = ?", userID, models.Expense).
		Count(&expenseCount)

	averageExpense := float64(0)
	if expenseCount > 0 {
		averageExpense = totalExpenses / float64(expenseCount)
	}

	// Generate AI-based recommendations
	recommendations := s.generateRecommendations(totalIncome, totalExpenses, savingsRate, categoryInsights, averageExpense)

	return &models.InsightResponse{
		TotalIncome:     totalIncome,
		TotalExpenses:   totalExpenses,
		NetIncome:       totalIncome - totalExpenses,
		TopCategories:   categoryInsights,
		MonthlyTrend:    monthlyTrends,
		Recommendations: recommendations,
		SavingsRate:     savingsRate,
		AverageExpense:  averageExpense,
	}, nil
}

func (s *InsightsService) getMonthlyTrends(userID uint) []models.MonthlyTrend {
	db := database.GetDB()
	var trends []models.MonthlyTrend

	// Get data for last 6 months
	for i := 5; i >= 0; i-- {
		month := time.Now().AddDate(0, -i, 0)
		monthStr := month.Format("2006-01")
		monthName := month.Format("Jan 2006")

		var income, expenses float64

		db.Model(&models.Transaction{}).
			Where("user_id = ? AND type = ? AND strftime('%Y-%m', date) = ?", userID, models.Income, monthStr).
			Select("COALESCE(SUM(amount), 0)").
			Scan(&income)

		db.Model(&models.Transaction{}).
			Where("user_id = ? AND type = ? AND strftime('%Y-%m', date) = ?", userID, models.Expense, monthStr).
			Select("COALESCE(SUM(amount), 0)").
			Scan(&expenses)

		trends = append(trends, models.MonthlyTrend{
			Month:    monthName,
			Income:   income,
			Expenses: expenses,
		})
	}

	return trends
}

func (s *InsightsService) generateRecommendations(totalIncome, totalExpenses, savingsRate float64, 
	categories []models.CategoryInsight, averageExpense float64) []string {
	
	var recommendations []string

	// Savings rate recommendations
	if savingsRate < 10 {
		recommendations = append(recommendations, "Your savings rate is below 10%. Consider reducing discretionary spending to increase savings.")
	} else if savingsRate > 30 {
		recommendations = append(recommendations, "Excellent savings rate! You're on track for strong financial health.")
	} else {
		recommendations = append(recommendations, "Good savings rate. Try to aim for 20-30% for optimal financial security.")
	}

	// Expense pattern analysis
	if len(categories) > 0 {
		topCategory := categories[0]
		if topCategory.Amount > totalExpenses*0.4 {
			recommendations = append(recommendations, 
				fmt.Sprintf("Your spending on '%s' represents %.1f%% of total expenses. Consider reviewing this category for potential savings.", 
				topCategory.Category, (topCategory.Amount/totalExpenses)*100))
		}
	}

	// Income vs Expenses analysis
	if totalExpenses > totalIncome {
		recommendations = append(recommendations, "You're spending more than you earn. Consider creating a budget to track and reduce expenses.")
	}

	// High average expense warning
	if averageExpense > 500 && len(categories) > 0 {
		recommendations = append(recommendations, "Your average expense per transaction is high. Review large purchases and consider if they align with your financial goals.")
	}

	// Emergency fund recommendation
	if totalIncome > 0 {
		monthlyExpenses := totalExpenses / 12
		if totalIncome-totalExpenses < monthlyExpenses*3 {
			recommendations = append(recommendations, "Consider building an emergency fund covering 3-6 months of expenses.")
		}
	}

	// Default recommendation if none generated
	if len(recommendations) == 0 {
		recommendations = append(recommendations, "Keep tracking your expenses and maintain good financial habits!")
	}

	return recommendations
}