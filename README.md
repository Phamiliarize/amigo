# amigo
amigo is a modern bulletin board software that aims to be performant, easily extensible, and simple to deploy.

# Features (Roadmap)
More than categories, threads, and posts, a modern bulletin board needs;

- __Auth__ (40%)
  - Route-based RBAC Authorization ✅
  - Authentication
    - Amigo follows a BYO model- it seamlessly integrates with your existing identity solution by using [JWTs](https://jwt.io/). Simply map Amigo roles to your users in your identity provider and pass the JWK URL your instance should use to verify tokens.
- __BBS__ (0%)
  - Progressive enhancement/SSR for SEO
  - Categories, threads, and posts
  - Mod Tools
    - Needs prep/user stories
  - Private Functionality
    - Private Categories
    - Private Threads?
  - User Orgs
- __Extensible Markdown__ (0%)
  - Safe, sanitized markdown that can be extended for your usecase.
- __Themes__ (70%)
  - Easily create themes via custom CSS; over-ride only the HTML templates you need
  - Users can select a theme as part of their preferences
- __Hooks__ (0%)
  - Extend functionality and execute your own scripts during the lifecycle of specific actions (tbd)
- __In-App Notifications__ (0%)
  - Notify users of replies; staff of new posts

~~Levels (0%) - Activity and kudos earn levels; different actions can be locked behind levels, giving a tool for filtering out bots, trolls, and bad actors.
Profiles (0%) - User's can assume a different profile and post as that profile. Profiles are extensible.~~
: