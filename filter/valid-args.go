// SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
// SPDX-FileContributor: Matthew Taylor (Bertie690)
//
// SPDX-License-Identifier: GPL-3.0-or-later

package filter

import (
	"regexp"
	"strings"

	"github.com/Bertie690/gh-pr-list/utils"
)

// All valid arguments that can be obtained from JSON.
var validArgs = utils.NewSet([]string{
	"additions",
	"assignees",
	"author",
	"autoMergeRequest",
	"baseRefName",
	"baseRefOid",
	"body",
	"changedFiles",
	"closed",
	"closedAt",
	"closingIssuesReferences",
	"comments",
	"commits",
	"createdAt",
	"deletions",
	"files",
	"fullDatabaseId",
	"headRefName",
	"headRefOid",
	"headRepository",
	"headRepositoryOwner",
	"id",
	"isCrossRepository",
	"isDraft",
	"labels",
	"latestReviews",
	"maintainerCanModify",
	"mergeCommit",
	"mergeStateStatus",
	"mergeable",
	"mergedAt",
	"mergedBy",
	"milestone",
	"number",
	"potentialMergeCommit",
	"reactionGroups",
	"reviewDecision",
	"reviewRequests",
	"reviews",
	"state",
	"statusCheckRollup",
	"title",
	"updatedAt",
	"url",
}...)

// Regex to extract field names from dot notation property access.
var fieldsRe = regexp.MustCompile(`\.(\w+)`)

// getRequiredFields returns a comma-separated string of all arguments required by either the filter or template.
func getRequiredFields(query, template string) string {
	var b strings.Builder
	seen := utils.NewSet[string]()

	// NB: We can use the same regex since both JQ and GO templates use dot notation to access fields
	queryMatch := fieldsRe.FindAllStringSubmatch(query, -1)
	templateMatch := fieldsRe.FindAllStringSubmatch(template, -1)

	checkMatches(queryMatch, seen, &b)
	checkMatches(templateMatch, seen, &b)

	result := b.String()
	if len(result) > 0 {
		return result[:len(result)-1] // Remove trailing comma
	}
	return ""
}

// checkMatches adds valid fields from matches to the builder if they haven't been seen before.
func checkMatches(matches [][]string, seen utils.Set[string], b *strings.Builder) {
	for _, match := range matches {
		if len(match) < 2 {
			continue
		}
		field := match[1]
		if validArgs.Contains(field) && !seen.Contains(field) {
			b.WriteString(field)
			b.WriteString(",")
			seen.Add(field)
		}
	}
}
