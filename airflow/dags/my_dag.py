from datetime import timedelta

import airflow
from airflow import DAG
from airflow.operators.python_operator import PythonOperator
from time import sleep

default_args = {
	"owner": "airflow",
	"start_date": airflow.utils.dates.days_ago(2),
	"depends_on_past": False,
	"retries": 2,
	"retry_delay": timedelta(seconds=10),
}

with DAG(
	"my_dag",
	default_args=default_args,
	description="My DAG",
	schedule_interval=timedelta(days=1),
) as dag:

	def taskn(time: int):
		print(f"sleeping for {time}")
		sleep(time)

	t1 = PythonOperator(
		task_id="my_task1",
		python_callable=taskn,
		op_kwargs={"time": 10}
	)

	t2 = PythonOperator(
		task_id="my_task2",
		python_callable=taskn,
		op_kwargs={"time": 20}
	)
	
	t25 = PythonOperator(
		task_id="my_task2.5",
		python_callable=taskn,
		op_kwargs={"time": 25}
	)

	t3 = PythonOperator(
		task_id="my_task3",
		python_callable=taskn,
		op_kwargs={"time": 30}
	)

	t4 = PythonOperator(
		task_id="my_task4",
		python_callable=taskn,
		op_kwargs={"time": 40}
	)

	t1 >> [t2, t25] >> t3 >> t4
