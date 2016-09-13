package gowyn

import (
	"fmt"
	"github.com/libgit2/git2go"
)

var statusOption = git.StatusOptions{
	git.StatusShowIndexAndWorkdir,
	git.StatusOptIncludeUntracked,
	[]string{},
}

func CheckStateOfGitObjects(group *string) {

	if *group == "" {

		gitRepositories, err := globalContainer.S(FILENAME_PATH).Children()

		if err != nil {
			ErrorTracer.Fatalf("Canno't retrieve git repositories from %s\n", FILENAME_PATH)
		}

		for _, child := range gitRepositories {
			checkStateOfGitObject(child.Data().(string))
		}

	} else {

		gitRepositories, err := globalContainer.S(GROUPS_PATH, *group).Children()

		if err != nil {
			ErrorTracer.Fatalf("Canno't retrieve the group %s from %s\n", *group, GROUPS_PATH)
		}

		for _, child := range gitRepositories {
			checkStateOfGitObject(child.Data().(string))
		}

	}

}

func checkStateOfGitObject(pathdir string) {

	gitRepository, err := git.OpenRepository(pathdir)

	if err != nil {
		ErrorTracer.Println("Cannot' open git repository (%s), due to \"%s\"", pathdir, err)
	}

	fmt.Printf("* %s:\n", gitRepository.Workdir())

	if index, err := gitRepository.Index(); err == nil {
		if index.HasConflicts() {
			fmt.Printf("\tWARNING::YOU HAVE SOME CONFLICTS!")
		}
	}

	if listOfUntrackedFiles, err := gitRepository.StatusList(&statusOption); err == nil {
		if count, err := listOfUntrackedFiles.EntryCount(); err == nil {
			fmt.Printf("\t%d untracked files!\n", count)
		}
	} else {
		fmt.Printf("\tNo untracked files!\n")
	}

}
