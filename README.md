# goShaBenchmark
In order to run this test you need to genetate two additional files (I didn't want to use git-lfs)<br/>
```dd if=/dev/zero of=6G  bs=6G  count=1```<br/>
```dd if=/dev/zero of=1G  bs=1G  count=1```<br/>
Copy those files into resources<br/>


Run:<br/>
```go test -bench=.```<br/>
