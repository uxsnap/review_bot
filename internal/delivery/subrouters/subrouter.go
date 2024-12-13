package subrouters

type SubrouterDeps struct {
	KvClient          KvClient
	UsersService      UsersService
	CategoriesService CategoriesService
	QuestionsService  QuestionsService
}
