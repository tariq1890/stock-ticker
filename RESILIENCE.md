# Resilience

This section go over the possible actions one could take to improve resilience, observability and scalability of `stock-ticker`

1. Adding a Horizontal Pod Autoscaler (HPA) could improve the scalability of `stock-ticker`. The HPA could automatically
provision more replicas to handle higher than average loads. `stock-ticker` is a stateless service, so HPA would fit well.
2. We can improve resilience and availability by using pod AntiAffinity. For eg:- We could set a pod AntiAffinity policy
to distribute the replicas of `stock-ticker` across multiple availability zones or another Kubernetes node topology of your choice.
3. We have used `prometheus-client-golang` to implement a simple HTTP metrics generator. Adding more metrics to cover non-functional
and functional requirements would improve our monitoring. Adding alerting rules on these metrics would help us be better positioned
to mitigate incidents
4. Running load tests wil help with identify performance bottlenecks and tuning of HPAs
5. `PodDisruptionBudgets` (PDBs) are another way to ensure fault tolerance and availability. PDBs ensure that a stipulated
minimum of replicas are available at any given moment
6. Using a shared HA cache that is shared by all the replicas would help improve performance and mitigate risks of being
rate-limited by the AlphaVantage API
