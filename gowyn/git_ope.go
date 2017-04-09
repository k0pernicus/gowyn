package gowyn

/*
#include <git2.h>
#include <git2/sys/repository.h>
*/
import "C"

import (
	"fmt"

	"gopkg.in/libgit2/git2go.v24"
)

var stateToString = map[git.RepositoryState]string{
	git.RepositoryStateNone:                 "None",
	git.RepositoryStateMerge:                "Merge",
	git.RepositoryStateRevert:               "Revert",
	git.RepositoryStateCherrypick:           "Cherrypick",
	git.RepositoryStateBisect:               "Bisect",
	git.RepositoryStateRebase:               "Rebase",
	git.RepositoryStateRebaseInteractive:    "Rebase Interactive",
	git.RepositoryStateRebaseMerge:          "Rebase Merge",
	git.RepositoryStateApplyMailbox:         "Apply Mailbox",
	git.RepositoryStateApplyMailboxOrRebase: "Apply Mailbox or Rebase",
}

var statusOption = git.StatusOptions{
	git.StatusShowIndexAndWorkdir,
	git.StatusOptIncludeUntracked,
	[]string{},
}

func CheckStateOfGitObjects(group *string, full *bool) {

	if *group == "" {

		gitRepositories, err := globalContainer.S(FILENAME_PATH).Children()

		if err != nil {
			ErrorTracer.Fatalf("Canno't retrieve git repositories from %s\n", FILENAME_PATH)
		}

		for _, child := range gitRepositories {
			checkStateOfGitObject(child.Data().(string), full)
		}

	} else {

		gitRepositories, err := globalContainer.S(GROUPS_PATH, *group).Children()

		if err != nil {
			ErrorTracer.Fatalf("Canno't retrieve the group %s from %s\n", *group, GROUPS_PATH)
		}

		for _, child := range gitRepositories {
			checkStateOfGitObject(child.Data().(string), full)
		}

	}

}

func checkStateOfGitObject(pathdir string, full *bool) {

	gitRepository, err := git.OpenRepository(pathdir)

	if err != nil {
		ErrorTracer.Printf("Cannot' open git repository (%s), due to \"%s\"\n", pathdir, err)
	}

	fmt.Printf("* %s:\n", gitRepository.Workdir())

	fmt.Printf("\tCurrent state: %s\n", stateToString[gitRepository.State()])

	if index, err := gitRepository.Index(); err == nil {
		if index.HasConflicts() {
			fmt.Printf("\tWARNING::YOU HAVE SOME CONFLICTS!\n")
		}
	}

	if listOfUntrackedFiles, err := gitRepository.StatusList(&statusOption); err == nil {
		if count, err := listOfUntrackedFiles.EntryCount(); err == nil {
			fmt.Printf("\t%d untracked files!\n", count)
		}
	} else {
		fmt.Printf("\tNo untracked files!\n")
	}

	if *full {

		headReference, err := gitRepository.Head()
		if err != nil {
			ErrorTracer.Printf("Canno't get the head of the repository %s, due to \"%s\"", pathdir, err)
		} else {
			headReferenceName := headReference.Name()
			headReferenceBranch, _ := headReference.Branch().Name()
			headReferenceTarget := headReference.Target().String()
			fmt.Printf("\tCurrent reference: %s\n", headReferenceName)
			fmt.Printf("\tCurrent branch: %s\n", headReferenceBranch)
			fmt.Printf("\tLast commit id: %s\n", headReferenceTarget)
		}

	}

}
