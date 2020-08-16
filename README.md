## install

```
go install .
```

## Sample

```
cd path/to/kubernetes/website
git remote add upstream https://github.com/kubernetes/website.git

export ISSUE_CREATOR_TOKEN=xxxxxxxx
export IC_OLD_REVISION=upstream/release-1.17
export IC_NEW_REVISION=fb6364d
export IC_FILEPATH=docs/concepts/workloads/pods/pod-lifecycle.md

git diff $IC_OLD_REVISION $IC_NEW_REVISION -- content/en/$IC_FILEPATH | issue-creator
https://github.com/kubernetes/website/issues/new?title=...&body=...
```
