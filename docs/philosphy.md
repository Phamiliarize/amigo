# Philosophy

## Hybrid
You'll note we make use of alpine.js + traditional server-side rendering.

The general idea is some actions are best as *quick* data transfers, and thus utilize client-side calls to a REST API.

For example clicking the "reading mode" button and saving that to your preferences for later. A minimal data transfer and FE-first response is best for UX here.

At the same time, other things, such as reading a thread or posts are better as SSE. The load times are generally better, more minimal, and this also is great for SEO which is extremely important for a community, as they need to gain traction and be easily found, along with any information they dessiminate.


## Future
We are considering HTMX too as an option, but that does remove "data only" transfers as is.