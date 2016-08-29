package gowyn

import (
	"fmt"
	"github.com/libgit2/git2go"
)

func CheckStateOfGitObjects() {

	gitRepositories, err := globalContainer.S(FILENAME_PATH).Children()

	if err != nil {
		ErrorTracer.Fatalf("Canno't retrieve git repositories from %s\n", FILENAME_PATH)
	}

	for _, child := range gitRepositories {
		checkStateOfGitObject(child.Data().(string))
	}

}

func checkStateOfGitObject(pathdir string) {

	gitRepository, err := git.OpenRepository(pathdir)

	if err != nil {
		ErrorTracer.Println("Cannot' open git repository (%s), due to \"%s\"", pathdir, err)
	}

	var statusOption = git.StatusOptions{
		git.StatusShowIndexAndWorkdir,
		git.StatusOptIncludeUntracked,
		[]string{},
	}

	fmt.Printf("*%s:\n", gitRepository.Workdir())
	if listOfUntrackedFiles, err := gitRepository.StatusList(&statusOption); err == nil {
		if count, err := listOfUntrackedFiles.EntryCount(); err == nil {
			fmt.Printf("\t%d untracked files!\n", count)
		}
	} else {
		fmt.Printf("\tNo untracked files!\n")
	}

}
