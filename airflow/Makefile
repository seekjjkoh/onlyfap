pull-airflow:
	curl -LfO 'https://airflow.apache.org/docs/apache-airflow/2.1.3/docker-compose.yaml'

init-airflow:
	docker-compose up airflow-init

run-airflow:
	docker-compose up

cleanup-airflow:
	docker-compose down --volumes --remove-orphans
