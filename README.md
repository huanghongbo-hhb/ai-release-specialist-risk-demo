# AI Release Specialist Risk Demo

This repository intentionally contains release-risk signals for testing Zadig AI Release Specialist workflows.

## Intended Signals

- Unit test report contains failed test cases.
- Test pass rate is below 90%.
- Source code contains a few simple patterns for code scan demos.
- `Dockerfile` builds a small HTTP service listening on port 8080.
- `k8s/demo.yaml` can be imported as a Zadig YAML service for full workflow demos.

## Suggested Zadig Flow

Use this repository in a workflow such as:

```text
Code Scan -> Build -> Test -> AI Release Specialist -> Deploy
```

For demos, make sure the test step can publish JUnit-style results and allow the AI Release Specialist step to run.
