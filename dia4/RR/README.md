# Rollout and Rollback

The term *roll-back has* been adopted (actually misappropriated) in the context of IT management to mean `undo of changes during system upgrades'. 

An Rollout, this strategy that slowly replaces previous versions of an application with new versions of an application by completely replacing the infrastructure on which the application is running.

## Create Service 
```
docker service create --name vote --replicas 5 --publish 5000:80 instavote/vote
```

It specifies 5 replicas for the service. Behind the hood, this means four tasks (a task runs a container) of the vote service are now running on the Swarm. When sending a request to the vote service, on port 5000, this one is load-balanced (by default in a round robin way) towards one of the four tasks.

## Rolling Update
```
docker service update --image instavote/vote:indent vote
```

Using the following options in the service configuration, we can customize the way the update is done:

-  *--update-parallelis*: number of tasks to update at the same time.

-  *--update-delay*: time to wait before updating the next batch of tasks.

If we now update the service with a new image, the rolling update will act differently. Let’s update the service with the instavote/vote:movies:
```
docker service update --image instavote/vote:movies vote
```

Update/F5 on browser.

## Rollback

Because Swarm knows the previous specification of the vote service, it’s possible to rollback to this one with the following command:
```
docker service rollback vote
```

As for the configuration of the rolling update, we could’ve used the following options in the service definition to specify the way we want the rollback to be done:

- *--rollback-parallelism*
- *--rollback-delay*

Update/F5 on browser


Execute: 
```
docker service create --name whoami --replicas 2 --update-failure-action rollback --update-delay 10s \
--update-monitor 10s --publish 8000:8000 lucj/whoami:1.0
```

*--update-failure-action* flag and give it the rollback value so that an unsuccessful update will trigger a rollback automatically. Also, the *--update-delay* and *--update-monitor* are used to provide additional timing for the *HEALTHCHECK*.

Update image *lucj/whoami:2.0*:
```
docker service update --image lucj/whoami:2.0 whoami
```
