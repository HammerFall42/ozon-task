run_in_mem:
	docker-compose build server_in_mem
	docker-compose up server_in_mem

run_in_db:
	docker-compose build server_in_db
	docker-compose up server_in_db