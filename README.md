## Sample

```
export IC_OLD_REVISION=upstream/release-1.15
export IC_NEW_REVISION=upstream/release-1.17

git diff $IC_OLD_REVISION $IC_NEW_REVISION content/en/$IC_FILEPATH | issue-creator
```
