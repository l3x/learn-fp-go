package interfaces

type LocalHandler interface {
	FileExists(fileName string) (fileExists bool, err error)
}

type LocalRepo struct {
	localHandlers map[string]LocalHandler
	localHandler  LocalHandler
}

type LocalFileSystemRepo LocalRepo

func NewLocalRepo(localHandlers map[string]LocalHandler) *LocalFileSystemRepo {
	localRepo := new(LocalFileSystemRepo)
	localRepo.localHandlers = localHandlers
	localRepo.localHandler = localHandlers["LocalFileSystemRepo"]
	return localRepo
}

func (repo *LocalFileSystemRepo) FileExists(fileName string) (fileExists bool, err error) {
	return repo.localHandler.FileExists(fileName)
}
