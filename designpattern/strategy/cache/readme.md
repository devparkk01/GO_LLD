let's build a in-memory cache having limited size. It has some specified 
maximum capacity. 

Whenever we reach a max capacity, we need to evict some entries from our cache.

There are algorithms to evict cache
1. LRU
2. FIFO
3. LIFO
4. LFU

The main idea is how we can decouple our `cache` class from these algorithms so that we can change the algorithm at run time. Also, the cache should not change when a new algorithm is being added. 

This is where strategy pattern comes into the picture. It suggests creating a family of the algorithm with each algorithm having its own class. Each of these classes follows the same interface, and this makes the algorithm interchangeable within the family. Let's say the common interface name is `evictionAlgo`. 

Now our main ``cache` class will embed the evictionAlgo interface. Instead of implementing all types of eviction algorithms in itself, our cache class will delegate the execution to the evictionAlgo interface. Since evictionAlgo is an interface, we can change the algorithm in run time to either LRU, FIFO, LFU without changing the cache class.