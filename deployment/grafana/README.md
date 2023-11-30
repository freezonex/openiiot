## Grafana Role
### Viewer
This role allows a user to view dashboards, but they cannot edit them. This is suitable for users who only need to see dashboard data without the ability to modify it.

### Editor
Users with this role can view and edit dashboards. This role is ideal for users who need to create and configure dashboards.

### Admin
This role applies to a specific Grafana dashboard or folder. Users with this role can manage the dashboard or folder, including adding or deleting it, managing permissions, and more.

### GrafanaAdmin
 (often referred to as "Super Admin" or "Grafana Server Admin"): This role provides complete administrative access to all aspects of the Grafana instance, including server settings, users, permissions, data sources, and plugins. It's the highest level of access and should be limited to a small number of trusted users.

## Environment Variable
### allow anonymous login
```shell
GF_AUTH_ANONYMOUS_ENABLED=true
GF_AUTH_ANONYMOUS_ORG_NAME=Main Org.
GF_AUTH_ANONYMOUS_ORG_ROLE=Editor
```

### set server root
```shell
GF_SERVER_ROOT_URL=/apps/freezonex-dashboard/grafana
```

### allow to serve Grafana from a subpath of domain
```shell
GF_SERVER_SERVE_FROM_SUB_PATH=true
```

### allows Grafana dashboards to be embedded in iframes on other websites
```shell
GF_SECURITY_ALLOW_EMBEDDING=true
```

### generic OAuth2 configuration
```shell
GF_AUTH_GENERIC_OAUTH_ENABLED=true
GF_AUTH_GENERIC_OAUTH_AUTO_LOGIN=true
GF_AUTH_GENERIC_OAUTH_CLIENT_ID=openiiot
GF_AUTH_GENERIC_OAUTH_CLIENT_SECRET=openiiot_secret
GF_AUTH_GENERIC_OAUTH_AUTH_URL=http://127.0.0.1:8081/login/oauth/authorize
GF_AUTH_GENERIC_OAUTH_TOKEN_URL=http://127.0.0.1:8081/login/oauth/token
GF_AUTH_GENERIC_OAUTH_API_URL=http://127.0.0.1:8081/userinfo
GF_AUTH_GENERIC_OAUTH_ROLE_ATTRIBUTE_PATH = contains(role[*], 'GrafanaAdmin') && 'GrafanaAdmin' || contains(role[*], 'Admin') && 'Admin' || contains(role[*], 'Editor') && 'Editor' || 'Viewer'
```