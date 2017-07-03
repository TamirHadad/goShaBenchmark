# goShaBenchmark
In order to run this test you need to genetate two additional files (I didn't want to use git-lfs)
dd if=/dev/zero of=6G  bs=6G  count=1
dd if=/dev/zero of=1G  bs=1G  count=1
Copy those files into resources


Run:
go test -bench=.
