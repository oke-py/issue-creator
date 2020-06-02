## Sample

```
export ISSUE_CREATOR_TOKEN=xxxxxxxx
export IC_OLD_REVISION=upstream/release-1.15
export IC_NEW_REVISION=upstream/release-1.17
export IC_FILEPATH=docs/concepts/workloads/pods/pod-overview.md

git diff $IC_OLD_REVISION $IC_NEW_REVISION content/en/$IC_FILEPATH | issue-creator
https://github.com/kubernetes/website/issues/new?title=...&body=...
```
