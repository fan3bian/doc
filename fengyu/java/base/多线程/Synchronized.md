The synchronized keyword can be used on different levels:
- Instance methods
- Static methods
- Code blocks


#### Synchronized Instance Methods
Java uses a monitor also known as monitor lock or intrinsic lock, to provide synchronization.These monitors are bound to an object, thus all synchronized blocks of the same object can have only one thread executing them at the same time.Instance methods are synchronized over the instance of the class owning the method. Which means **only one thread per instance of the class can execute this methoed**.

#### Synchronized Static Methods
These methods are synchronized on the Class object associated with the class and since only one Class object exists per JVM per class, only one thread can execute inside a static synchronized method per class, irrespective of the number of instances it has.