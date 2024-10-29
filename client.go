// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package czlai

import (
	"context"
	"net/http"
	"os"

	"github.com/CZL-AI/czlai-go/internal/requestconfig"
	"github.com/CZL-AI/czlai-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the czlai API. You should not instantiate this client directly,
// and instead use the [NewClient] method instead.
type Client struct {
	Options            []option.RequestOption
	Aidoc              *AidocService
	AidocExotic        *AidocExoticService
	SessionRecords     *SessionRecordService
	MedicalRecords     *MedicalRecordService
	AICheckup          *AICheckupService
	AIConv             *AIConvService
	Users              *UserService
	Upload             *UploadService
	UploadImage        *UploadImageService
	UploadImageOss     *UploadImageOssService
	PetProfiles        *PetProfileService
	AIMemes            *AIMemeService
	MedicationAnalysis *MedicationAnalysisService
	Aipic              *AipicService
	Aipics             *AipicService
	AipicExotics       *AipicExoticService
	AITrials           *AITrialService
	AITrial            *AITrialService
	Policies           *PolicyService
	MagnumKeys         *MagnumKeyService
	Buriedpoints       *BuriedpointService
	Whitelist          *WhitelistService
	Pets               *PetService
	UserModuleUsages   *UserModuleUsageService
	Logins             *LoginService
	UserPoints         *UserPointService
	PointPrices        *PointPriceService
	PointDetails       *PointDetailService
	PointTasks         *PointTaskService
	CheckPoints        *CheckPointService
	UserAdvices        *UserAdviceService
	Evaluation         *EvaluationService
}

// NewClient generates a new client with the default option read from the
// environment (BEARER_TOKEN, BASIC_AUTH_USERNAME, BASIC_AUTH_PASSWORD). The option
// passed in as arguments are applied after these default arguments, and all option
// will be passed down to the services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("BEARER_TOKEN"); ok {
		defaults = append(defaults, option.WithBearerToken(o))
	}
	if o, ok := os.LookupEnv("BASIC_AUTH_USERNAME"); ok {
		defaults = append(defaults, option.WithUsername(o))
	}
	if o, ok := os.LookupEnv("BASIC_AUTH_PASSWORD"); ok {
		defaults = append(defaults, option.WithPassword(o))
	}
	opts = append(defaults, opts...)

	r = &Client{Options: opts}

	r.Aidoc = NewAidocService(opts...)
	r.AidocExotic = NewAidocExoticService(opts...)
	r.SessionRecords = NewSessionRecordService(opts...)
	r.MedicalRecords = NewMedicalRecordService(opts...)
	r.AICheckup = NewAICheckupService(opts...)
	r.AIConv = NewAIConvService(opts...)
	r.Users = NewUserService(opts...)
	r.Upload = NewUploadService(opts...)
	r.UploadImage = NewUploadImageService(opts...)
	r.UploadImageOss = NewUploadImageOssService(opts...)
	r.PetProfiles = NewPetProfileService(opts...)
	r.AIMemes = NewAIMemeService(opts...)
	r.MedicationAnalysis = NewMedicationAnalysisService(opts...)
	r.Aipic = NewAipicService(opts...)
	r.Aipics = NewAipicService(opts...)
	r.AipicExotics = NewAipicExoticService(opts...)
	r.AITrials = NewAITrialService(opts...)
	r.AITrial = NewAITrialService(opts...)
	r.Policies = NewPolicyService(opts...)
	r.MagnumKeys = NewMagnumKeyService(opts...)
	r.Buriedpoints = NewBuriedpointService(opts...)
	r.Whitelist = NewWhitelistService(opts...)
	r.Pets = NewPetService(opts...)
	r.UserModuleUsages = NewUserModuleUsageService(opts...)
	r.Logins = NewLoginService(opts...)
	r.UserPoints = NewUserPointService(opts...)
	r.PointPrices = NewPointPriceService(opts...)
	r.PointDetails = NewPointDetailService(opts...)
	r.PointTasks = NewPointTaskService(opts...)
	r.CheckPoints = NewCheckPointService(opts...)
	r.UserAdvices = NewUserAdviceService(opts...)
	r.Evaluation = NewEvaluationService(opts...)

	return
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// If a byte slice or an [io.Reader] is supplied to params, it will be used as-is
// for the request body.
//
// The params is by default serialized into the body using [encoding/json]. If your
// type implements a MarshalJSON function, it will be used instead to serialize the
// request. If a URLQuery method is implemented, the returned [url.Values] will be
// used as query strings to the url.
//
// If your params struct uses [param.Field], you must provide either [MarshalJSON],
// [URLQuery], and/or [MarshalForm] functions. It is undefined behavior to use a
// struct uses [param.Field] without specifying how it is serialized.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url.
//
// The response body will be deserialized into the res variable, depending on its
// type:
//
//   - A pointer to a [*http.Response] is populated by the raw response.
//   - A pointer to a byte array will be populated with the contents of the request
//     body.
//   - A pointer to any other type uses this library's default JSON decoding, which
//     respects UnmarshalJSON if it is defined on the type.
//   - A nil value will not read the response body.
//
// For even greater flexibility, see [option.WithResponseInto] and
// [option.WithResponseBodyInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	opts = append(r.Options, opts...)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}
