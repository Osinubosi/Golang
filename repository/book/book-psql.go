package bookRepository

import (
	"book-list-app/models"
	"database/sql"
)

type BookRepository struct{}

func logFatal(err error) {
	if err != nil {
		logFatal(err)
	}
}

func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) ([]models.Book, error) {
	rows, err := db.Query("SELECT *FROM BOOKS")

	if err != nil {
		return []models.Book{}, err
	}

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)

		books = append(books, book)

	}

	if err != nil {
		return []models.Book{}, err
	}

	return books, nil
}

func (b BookRepository) GetBook(db *sql.DB, book models.Book, id int) (models.Book, error) {
	rows := db.QueryRow("SELECT * FROM BOOKS WHERE ID=$1", id)

	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (b BookRepository) AddBook(db *sql.DB, book models.Book) (int, error) {
	err := db.QueryRow("INSERT INTO BOOKS (TITLE,AUTHOR,YEAR) VALUES ($1,$2,$3) RETURNING ID;",
		book.Title, book.Author, book.Year).Scan(&book.ID)

	logFatal(err)

	if err != nil {
		return 0, err
	}

	return book.ID, nil
}

func (b BookRepository) UpdateBook(db *sql.DB, book models.Book) (int64, error) {
	result, err := db.Exec("UPDATE BOOKS SET TITLE=$1,AUTHOR =$2,YEAR=$3 WHERE ID=$4 RETURNING ID;",
		&book.Title, &book.Author, &book.Year, &book.ID)

	if err != nil {
		return 0, err
	}

	rowsUpdated, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsUpdated, nil
}

func (b BookRepository) RemoveBook(db *sql.DB, id int) (int64, error) {
	result, err := db.Exec("DELETE FROM BOOKS WHERE ID=$1", id)

	if err != nil {
		return 0, err
	}
	rowsDeleted, err := result.RowsAffected()
	logFatal(err)

	if err != nil {
		return 0, err
	}

	return rowsDeleted, nil
}
