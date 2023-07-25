package quotes

import (
	"errors"
	"go-server-app/internal/models"
	"go-server-app/internal/services"
)

type QuotesService struct {
	cache map[int]*[]models.Quote
}

func New() services.QuotesInterface {
	return &QuotesService{cache: map[int]*[]models.Quote{}}
}
func (qs *QuotesService) CreateQuote(empId int, quote *models.Quote) (*models.Quote, error) {
	if quote.Quote == "" || quote.Title == "" {
		return nil, errors.New("fields Cannot be Empty")
	}
	if quotesList, ok := qs.cache[empId]; ok {
		// If the employee already has quotes, append the new quote to the existing list
		*quotesList = append(*quotesList, *quote)
	} else {
		// If the employee does not have any quotes yet, create a new list with the quote
		qs.cache[empId] = &[]models.Quote{*quote}
	}

	// Return the newly added quote
	return quote, nil
}

func (qs *QuotesService) GetQuoteById(empId, quoteId int) (*models.Quote, error) {
	quotesList, ok := qs.cache[empId]
	if !ok {
		return nil, errors.New("employee not found")
	}
	for _, quote := range *quotesList {
		if quote.Id == quoteId {
			return &quote, nil
		}
	}

	return nil, errors.New("quote not found")
}

func (qs *QuotesService) GetAllQuotes() ([]models.Quote, error) {
	var allQuotes []models.Quote

	for _, quotesList := range qs.cache {
		allQuotes = append(allQuotes, *quotesList...)
	}

	return allQuotes, nil
}

func (qs *QuotesService) GetAllQuoteByEmpId(empId int) ([]models.Quote, error) {
	quotesList, ok := qs.cache[empId]
	if !ok {
		return nil, errors.New("employee not found")
	}

	return *quotesList, nil
}

func (qs *QuotesService) UpdateQuote(empId, quoteId int, updates map[string]string) (*models.Quote, error) {
	quotesList, ok := qs.cache[empId]
	if !ok {
		return nil, errors.New("employee not found")
	}

	for index, quote := range *quotesList {
		if quote.Id == quoteId {
			// Update the fields if they exist in the updates map
			if quoteText, exists := updates["quote"]; exists {
				(*quotesList)[index].Quote = quoteText
			}
			if title, exists := updates["title"]; exists {
				(*quotesList)[index].Title = title
			}

			return &(*quotesList)[index], nil
		}
	}

	return nil, errors.New("quote not found")
}

func (qs *QuotesService) DeleteQuote(empId, quoteId int) error {
	quotesList, ok := qs.cache[empId]
	if !ok {
		return errors.New("employee not found")
	}

	for index, quote := range *quotesList {
		if quote.Id == quoteId {
			// Delete the quote from the list
			*quotesList = append((*quotesList)[:index], (*quotesList)[index+1:]...)
			return nil
		}
	}

	return errors.New("quote not found")
}
