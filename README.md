# LearningGo
This repository is created for myself to practise go lang using some scratch projects

## Local Docker Push for Render (M1 Safe)

Use the helper script to build and push a linux/amd64 image from Apple Silicon.

1. Log in once:
	docker login
2. Run:
	./scripts/push-image-render.sh your-dockerhub-username

What it does:
- Builds and pushes your image as linux/amd64 for Render compatibility.
- Pushes two tags:
  - latest
  - git short SHA

Example output image names:
- your-dockerhub-username/go-api-app:latest
- your-dockerhub-username/go-api-app:abc1234

If your Render service is configured to pull from Docker Hub, point it to:
- your-dockerhub-username/go-api-app:latest

Or pin a release using a SHA tag.
