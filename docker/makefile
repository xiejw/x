DOCKERFILE_DIR  = dockerfiles

# NOTE: If $@ is xiejw/abc, then $(@F) is abc. For example,
# - xiejw/debian
# - xiejw/ubuntu
# - xiejw/clang
#
# NOTE: set Docker context as `assets` to avoid large file transfers.

xiejw/%:
	docker build -t $@ -f ${DOCKERFILE_DIR}/dockerfile.$(@F) assets

