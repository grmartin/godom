# Methodology #

## Test-Driven Development ##

We are experimenting with [Test-Driven Development](http://en.wikipedia.org/wiki/Test-driven_development) for godom.  This means we write the tests (and stub code) that fail first, then we write code to make them pass.  In Test-Driven Development, the test suite forms the basis of the requirements of the project.

The reason we feel this is a good approach for godom is simple: the DOM is an API that is fairly well-specified and broadly implemented.  This gives us a stable specification that we have to meet and a large number of interoperable implementations to verify our tests in.

In addition, as we build up the test suite, it will help to protect against regressions.  The more tests we have, the more protected we are and the more confidence we have that we are not breaking other things as we add features and fix bugs.

## Strategy ##

How do we pick which features to implement?  In what order?

In these beginning stages we are focused on getting a minimally useful subset of the DOM Core Level 1 implemented.  This means being able to do simple things like: traversing the tree, getting/setting attributes, getting node names/values, adding/removing nodes.

Once this minimally useful subset has been implemented, we will start exploring further aspects (XML namespaces, Element Traversal, Selectors API).