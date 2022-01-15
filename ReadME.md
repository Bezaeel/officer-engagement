# Officer-Engagement Service

This intends to be a service is part of a claims-systems, it deals with officers engagement and rotation.


# rotation
Get officers roll-call

# tag
Assign an officer to a claim case, 


# implementation
officers are assigned in a rotational manner

```
var officers = ["talabi", "ope", "jb", "sao];
officers[0].isTagged = true;
```
on next call

```
officers[0].isTagged = false;
officers[1].isTagged = true;
```

rotation becomes

```
var officers = ["ope", "jb", "sao, "talabi"];
```