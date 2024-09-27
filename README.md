Along with CURD operations I have also added a dir for learning go routines and how channels work. See the demo output run below 

// Demo run without using go routine
```
// Sample output :
// This is the data for: Delhi {Main:{Temp:303.2}}
// This is the data for: Bengaluru {Main:{Temp:297.38}}
// This is the data for: Mumbai {Main:{Temp:298.14}}
// This is the data for: Patna {Main:{Temp:301.11}}
// This operation took :  559.819125ms
```

// Demo run with go routines
```
// This is the city Patna
// This is the city Delhi
// This is the city Banglore
// This is the city Mumbai
// This operation took :  157.35ms
```
The code is nearly 3x faster with go routines. Happy Coding !!
