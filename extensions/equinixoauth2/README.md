# Equinix OAuth2 Extension

This package provides a custom, context-aware TokenSource and HTTP Transport that can be used to automatically exchange an Equinix client ID and secret for an OAuth token before making requests to Equinix APIs.

## Usage

1. Import the extension

    ```
    import "github.com/equinix/equinix-sdk-go/extensions/equinixoauth2"
    ```

2. Prepare configuration and create authenticating transport 

    ```
    authConfig := equinixoauth2.Config{
        ClientID:     "myClientId",
        ClientSecret: "myClientSecret"
        BaseURL:      "https://api.equinix.com"}

    //*http.RoundTripper is returned
    authTransport := authConfig.New()
    ```

3. Use authenticating transport for requests to an Equinix API

    The `*http.RoundTripper` created by the OAuth2 extension will handle obtaining and refreshing a token as well as adding a properly-formatted Authorization header to API requests. 

    The sample code below demonstrates how to use the OAuth2 extension with the [Fabric API](../../services/fabricv4).
    
    ```
    configuration := fabricv4.NewConfiguration()
    configuration.HTTPClient = &http.Client{
        Transport: authTransport
    }
    client := fabricv4.NewAPIClient(configuration)
    conn, resp, err := client.ConnectionsApi.GetConnectionByUuid(ctx, "<my_connection_uuid>").Execute()
    ```

    You can run the [fabricv4/get_connection example](../../examples/services/fabricv4/get_connection) to see this code in action.