package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	DefaultUsersTable               = "user"
	DefaultProblemsTable            = "problem"
	DefaultSubmissionsTable         = "submission"
	DefaultTestCasesTable           = "test_case"
	DefaultCompetitionsTable        = "competition"
	DefaultCompetitionProblemsTable = "competition_problem"
	DefaultCompetitionUsersTable    = "competition_user"
)

type Tables struct {
	Users               string
	Problems            string
	Submissions         string
	TestCases           string
	Competitions        string
	CompetitionProblems string
	CompetitionUsers    string
}

type Database struct {
	NextJudgeDB *sql.DB
	TableNames  Tables
}

type NextJudgeDB interface {
	GetUsers() ([]User, error)
}

func NewDatabase() (*Database, error) {
	var err error
	db := &Database{
		TableNames: Tables{
			Users:               cfg.UsersTable,
			Problems:            cfg.ProblemsTable,
			Submissions:         cfg.SubmissionsTable,
			TestCases:           cfg.TestCasesTable,
			Competitions:        cfg.CompetitionsTable,
			CompetitionProblems: cfg.CompetitionProblemsTable,
			CompetitionUsers:    cfg.CompetitionUsersTable,
		},
	}
	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, strconv.FormatInt(cfg.Port, 64), cfg.Username, cfg.Password, cfg.DBName)
	db.NextJudgeDB, err = sql.Open(cfg.DBDriver, dataSource)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (d Database) GetUser() ([]User, error) {
	// Get users
	return nil, nil
}
