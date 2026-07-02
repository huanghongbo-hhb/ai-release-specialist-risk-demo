FROM scratch

COPY ai-release-specialist-risk-demo /ai-release-specialist-risk-demo
EXPOSE 8080
ENTRYPOINT ["/ai-release-specialist-risk-demo"]
