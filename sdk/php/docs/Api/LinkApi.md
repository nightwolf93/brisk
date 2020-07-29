# Swagger\Client\LinkApi

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**apiV1LinkPut**](LinkApi.md#apiv1linkput) | **PUT** /api/v1/link | Create new link
[**slugGet**](LinkApi.md#slugget) | **GET** /:slug | Visit a link

# **apiV1LinkPut**
> \Swagger\Client\Model\CreateLinkResponse apiV1LinkPut($url, $ttl, $x_client_id, $x_client_secret)

Create new link

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

$apiInstance = new Swagger\Client\Api\LinkApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client()
);
$url = "url_example"; // string | 
$ttl = 56; // int | 
$x_client_id = "x_client_id_example"; // string | 
$x_client_secret = "x_client_secret_example"; // string | 

try {
    $result = $apiInstance->apiV1LinkPut($url, $ttl, $x_client_id, $x_client_secret);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LinkApi->apiV1LinkPut: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **string**|  |
 **ttl** | **int**|  |
 **x_client_id** | **string**|  | [optional]
 **x_client_secret** | **string**|  | [optional]

### Return type

[**\Swagger\Client\Model\CreateLinkResponse**](../Model/CreateLinkResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **slugGet**
> slugGet()

Visit a link

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

$apiInstance = new Swagger\Client\Api\LinkApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client()
);

try {
    $apiInstance->slugGet();
} catch (Exception $e) {
    echo 'Exception when calling LinkApi->slugGet: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

