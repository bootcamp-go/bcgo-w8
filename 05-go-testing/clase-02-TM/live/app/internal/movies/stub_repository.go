package movies

// DummyRepository is a dummy movie repository.
// type DummyRepository struct {}

// func (rp *DummyRepository) GetMovie(id int) (mv Movie, err error) {
// 	return
// }

// ------------------------------------------------------------
// constructor
func NewRepositoryStubFunc() *RepositoryStubFunc {
	return &RepositoryStubFunc{}
}

// RepositoryStubFunc is a function that returns a movie repository stub.
type RepositoryStubFunc struct {
	GetMovieStub func(id int) (mv Movie, err error)
	// Err error
}

func (rp *RepositoryStubFunc) GetMovie(id int) (mv Movie, err error) {
	// if rp.Err != nil {
	// 	err = rp.Err
	// 	return
	// }
	// mv = Movie{
	// 	ID:          1,
	// 	Title:       "The Matrix",
	// 	ReleaseYear: 1999,
	// }
	mv, err = rp.GetMovieStub(id)
	return
}


// ------------------------------------------------------------
// constructor
func NewRepositoryStubReturn() *RepositoryStubReturn {
	return &RepositoryStubReturn{}
}

// RepositoryStubReturn is a movie repository stub that returns a movie.
type RepositoryStubReturn struct {
	GetMovieStub struct {mv Movie; err error}
}

func (rp *RepositoryStubReturn) GetMovie(id int) (mv Movie, err error) {
	mv = rp.GetMovieStub.mv
	err = rp.GetMovieStub.err
	return
}



// ------------------------------------------------------------
// constructor
func NewRepositoryStubArgs() *ServiceStubArgs {
	return &ServiceStubArgs{}
}

// ServiceStubArgs is a stub implementation of the Service interface.
type ServiceStubArgs struct {
	GetMovieArgs   Args
}
// GetMovie calls the GetMovieArgs.
func (s *ServiceStubArgs) GetMovie(id int) (mv Movie, err error) {
	return s.GetMovieArgs.Get(0).(Movie), s.GetMovieArgs.Error(1)
}


// ------------------------------------------------------------
func NewArgs(a ...any) Args {
	return Args{a}
}

// Args is an slice of arguments of type any.
type Args struct {
	a []any
}
// get returns the argument at the given index.
func (a *Args) Get(i int) any {
	return a.a[i]
}
func (a *Args) String(i int) string {
	return a.a[i].(string)
}
func (a *Args) Int(i int) int {
	return a.a[i].(int)
}
func (a *Args) Float64(i int) float64 {
	return a.a[i].(float64)
}
func (a *Args) Bool(i int) bool {
	return a.a[i].(bool)
}
// pointer cases
func (a *Args) Error(i int) error {
	var e error
	if a.a[i] != nil {
		e = a.a[i].(error)
	}
	return e
}
