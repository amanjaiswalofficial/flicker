# flicker

## Roadmap
### Components of framework
1. Cluster Manager: A simple cluster manager that can discover and manage nodes in the cluster. This could be done using a simple peer-to-peer network or a centralized configuration file.

2. Job Scheduler: A job scheduler that can schedule and assign tasks to nodes in the cluster. The scheduler should be able to handle dependencies between tasks and ensure that tasks are executed in the correct order.

3. Communication Protocol: A communication protocol for nodes in the cluster to exchange messages and coordinate task execution. This could be done using simple message passing or RPC.

4. Task Executor: A task executor that can execute tasks on nodes in the cluster. The executor should be able to communicate with the scheduler and report task status back to the scheduler.

5. Data Partitioning: A simple data partitioning mechanism that can partition input data into smaller, manageable chunks. This could be done using simple file partitioning or record-based partitioning.

6. Fault Tolerance: A basic fault tolerance mechanism that can handle node failures and task failures. This could be done using simple checkpointing and recovery mechanisms


### Types of components
1. Job submitter: This component can be an API that accepts job submissions from users. The API can be implemented using a web framework such as Flask or Django in Python.

2. Job scheduler: This component can be a standalone service that manages the scheduling and execution of jobs submitted by users. It can be implemented using a messaging queue such as RabbitMQ or Apache Kafka to receive job requests from the job submitter, and then dispatch them to the available executors. The scheduler can be implemented using a programming language such as Go or Rust.

3. Executor: This component can be a script that runs on each worker node, and performs the actual computation of the jobs. The executor can be implemented using a programming language such as Python or Go, and can use a library such as Dask or Apache Spark to parallelize the computation.

4. BTS service: This component can be implemented as a standalone service that facilitates data transfer between nodes during a shuffle operation. It can be implemented using a programming language such as Go or Rust, and can use a network protocol such as TCP or UDP to transfer data.

5. Data storage: This component can be a storage system that holds the input and output data of the jobs. The storage system can be implemented using a database such as PostgreSQL or a distributed filesystem such as Hadoop HDFS


### Sample framework project
1. Job Submission:

    Submit a job to the job scheduler by passing the GCS file location as input.
    The job scheduler receives the job submission and creates a DAG of tasks to process the data.
    The job scheduler assigns the tasks to available executors.

2. Reading Data:

    Each executor reads a partition of the input file from the GCS location assigned to it by the job scheduler.
    The executor processes the partition to generate intermediate results.
    The intermediate results are passed to the next stage of the processing pipeline.

3. GroupBy and Shuffle:

    The job scheduler identifies the stage in the DAG where the GroupBy operation is performed.
    The job scheduler creates a shuffle map that identifies the partitions of data that need to be exchanged between executors.
    The job scheduler assigns BTS services to move data between executors according to the shuffle map.
    The executors exchange intermediate results with each other via BTS service.

4. Aggregate Results:

    Each executor receives shuffled data from other executors.
    The executor performs the Max operation on the received data and generates a final intermediate result.
    The intermediate result is passed to the next stage of the processing pipeline.

5. Write to GCS:

    The job scheduler identifies the stage in the DAG where the final output is generated.
    The job scheduler assigns a single executor to write the final output to the GCS location.
    The executor writes the final output to the GCS location.

6. Job Completion:

    The job scheduler receives completion messages from all the executors.
    The job scheduler aggregates the completion messages to determine if the job completed successfully.
    If the job completed successfully, the job scheduler notifies the user and returns the final output
