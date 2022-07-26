// Code generated by entc, DO NOT EDIT.

package db

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/repository"
)

// Repository is the model entity for the Repository schema.
type Repository struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// RepositoryName holds the value of the "repository_name" field.
	RepositoryName string `json:"repository_name,omitempty"`
	// GitOrganization holds the value of the "git_organization" field.
	GitOrganization string `json:"git_organization,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// GitURL holds the value of the "git_url" field.
	GitURL string `json:"git_url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RepositoryQuery when eager-loading is set.
	Edges RepositoryEdges `json:"edges"`
}

// RepositoryEdges holds the relations/edges for other nodes in the graph.
type RepositoryEdges struct {
	// Workflows holds the value of the workflows edge.
	Workflows []*Workflows `json:"workflows,omitempty"`
	// Codecov holds the value of the codecov edge.
	Codecov []*CodeCov `json:"codecov,omitempty"`
	// ProwSuites holds the value of the prow_suites edge.
	ProwSuites []*ProwSuites `json:"prow_suites,omitempty"`
	// ProwJobs holds the value of the prow_jobs edge.
	ProwJobs []*ProwJobs `json:"prow_jobs,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// WorkflowsOrErr returns the Workflows value or an error if the edge
// was not loaded in eager-loading.
func (e RepositoryEdges) WorkflowsOrErr() ([]*Workflows, error) {
	if e.loadedTypes[0] {
		return e.Workflows, nil
	}
	return nil, &NotLoadedError{edge: "workflows"}
}

// CodecovOrErr returns the Codecov value or an error if the edge
// was not loaded in eager-loading.
func (e RepositoryEdges) CodecovOrErr() ([]*CodeCov, error) {
	if e.loadedTypes[1] {
		return e.Codecov, nil
	}
	return nil, &NotLoadedError{edge: "codecov"}
}

// ProwSuitesOrErr returns the ProwSuites value or an error if the edge
// was not loaded in eager-loading.
func (e RepositoryEdges) ProwSuitesOrErr() ([]*ProwSuites, error) {
	if e.loadedTypes[2] {
		return e.ProwSuites, nil
	}
	return nil, &NotLoadedError{edge: "prow_suites"}
}

// ProwJobsOrErr returns the ProwJobs value or an error if the edge
// was not loaded in eager-loading.
func (e RepositoryEdges) ProwJobsOrErr() ([]*ProwJobs, error) {
	if e.loadedTypes[3] {
		return e.ProwJobs, nil
	}
	return nil, &NotLoadedError{edge: "prow_jobs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Repository) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case repository.FieldRepositoryName, repository.FieldGitOrganization, repository.FieldDescription, repository.FieldGitURL:
			values[i] = new(sql.NullString)
		case repository.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Repository", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Repository fields.
func (r *Repository) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case repository.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				r.ID = *value
			}
		case repository.FieldRepositoryName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field repository_name", values[i])
			} else if value.Valid {
				r.RepositoryName = value.String
			}
		case repository.FieldGitOrganization:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field git_organization", values[i])
			} else if value.Valid {
				r.GitOrganization = value.String
			}
		case repository.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				r.Description = value.String
			}
		case repository.FieldGitURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field git_url", values[i])
			} else if value.Valid {
				r.GitURL = value.String
			}
		}
	}
	return nil
}

// QueryWorkflows queries the "workflows" edge of the Repository entity.
func (r *Repository) QueryWorkflows() *WorkflowsQuery {
	return (&RepositoryClient{config: r.config}).QueryWorkflows(r)
}

// QueryCodecov queries the "codecov" edge of the Repository entity.
func (r *Repository) QueryCodecov() *CodeCovQuery {
	return (&RepositoryClient{config: r.config}).QueryCodecov(r)
}

// QueryProwSuites queries the "prow_suites" edge of the Repository entity.
func (r *Repository) QueryProwSuites() *ProwSuitesQuery {
	return (&RepositoryClient{config: r.config}).QueryProwSuites(r)
}

// QueryProwJobs queries the "prow_jobs" edge of the Repository entity.
func (r *Repository) QueryProwJobs() *ProwJobsQuery {
	return (&RepositoryClient{config: r.config}).QueryProwJobs(r)
}

// Update returns a builder for updating this Repository.
// Note that you need to call Repository.Unwrap() before calling this method if this Repository
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Repository) Update() *RepositoryUpdateOne {
	return (&RepositoryClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Repository entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Repository) Unwrap() *Repository {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("db: Repository is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Repository) String() string {
	var builder strings.Builder
	builder.WriteString("Repository(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", repository_name=")
	builder.WriteString(r.RepositoryName)
	builder.WriteString(", git_organization=")
	builder.WriteString(r.GitOrganization)
	builder.WriteString(", description=")
	builder.WriteString(r.Description)
	builder.WriteString(", git_url=")
	builder.WriteString(r.GitURL)
	builder.WriteByte(')')
	return builder.String()
}

// Repositories is a parsable slice of Repository.
type Repositories []*Repository

func (r Repositories) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
