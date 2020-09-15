![](assets/title.png)

## About

Personal project to create random data.

Target to produce 100,000+ unique user profiles and addresses in <2 seconds

```
§ serendipity person -s seed123 -n 100000 | jq '.[].preferred_username' | sort | uniq -c | sort -nr > duplicates.txt
> 1.68s user 0.27s system 114% cpu 1.702 total
```
