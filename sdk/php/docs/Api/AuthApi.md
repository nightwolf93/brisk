# Swagger\Client\AuthApi

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**apiV1CredentialPut**](AuthApi.md#apiv1credentialput) | **PUT** /api/v1/credential | Create new credential

# **apiV1CredentialPut**
> apiV1CredentialPut($client_id, $client_secret, $x_client_id, $x_client_secret)

Create new credential

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

$apiInstance = new Swagger\Client\Api\AuthApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client()
);
$client_id = "client_id_example"; // string | 
$client_secret = "client_secret_example"; // string | 
$x_client_id = "x_client_id_example"; // string | 
$x_client_secret = "x_client_secret_example"; // string | 

try {
    $apiInstance->apiV1CredentialPut($client_id, $client_secret, $x_client_id, $x_client_secret);
} catch (Exception $e) {
    echo 'Exception when calling AuthApi->apiV1CredentialPut: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **client_id** | **string**|  |
 **client_secret** | **string**|  |
 **x_client_id** | **string**|  | [optional]
 **x_client_secret** | **string**|  | [optional]

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

