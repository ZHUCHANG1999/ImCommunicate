user-rpc-test:
	@make -f deploy/mk/user-rpc.mk release-test

release-test: user-rpc-test user-api-test social-rpc-test social-api-test

install-server:
	cd ./deploy/script && chmod +x release-test.sh && ./release-test.sh
