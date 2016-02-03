[← Integrate Other Repos](integarte-other-repos.md)

# Add Additional Dimensions

You can also add additional Dimensions to the integration test.

Let's update our `docker-compose.yml` file to include an additional Dimension:

```
# ...
- XLANG_DIMENSION_BEHAVIOR=dance,run
- XLANG_DIMENSION_SPEED=fast,slow
```

Now run xlang:

```
$ docker-compose run xlang

Beginning matrix of tests...

  STATUS  |  CLIENT   |      RESPONSE       | SPEED | BEHAVIOR
+---------+-----------+---------------------+-------+----------+
  PASSED  | client    | ok                  | fast  | dance
  SKIPPED | client    | 404                 | fast  | run
  PASSED  | client    | ok                  | slow  | dance
  SKIPPED | client    | 404                 | slow  | run
  PASSED  | newclient | ok                  | fast  | dance
  PASSED  | newclient | ok                  | fast  | run
  PASSED  | newclient | ok                  | slow  | dance
  PASSED  | newclient | ok                  | slow  | run

```

Awesome - a test was triggered for every combination of "speed" and "behavior".

In this way, we can develop our Test Client's to address multiple dimensions of integration.

[Run During Continuous Integartion →](add-to-ci.md)
