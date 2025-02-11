package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	dbmocks "github.com/github-real-lb/bookings-web-app/db/mocks"
	loggermocks "github.com/github-real-lb/bookings-web-app/util/loggers/mocks"
	"github.com/github-real-lb/bookings-web-app/util/mailers"
	mailermocks "github.com/github-real-lb/bookings-web-app/util/mailers/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type TestServer struct {
	*Server
	MockDBStore     *dbmocks.MockDBStore
	MockErrorLogger *loggermocks.MockLogger
	MockInfoLogger  *loggermocks.MockLogger
	MockMailer      *mailermocks.MockMailer
}

// NewTestServer creates and returns a test server connected to a mock database store
func NewTestServer(t *testing.T) *TestServer {
	// create mocks
	mockDBStore := dbmocks.NewMockDBStore(t)
	mockErrorLogger := loggermocks.NewMockLogger(t)
	mockInfoLogger := loggermocks.NewMockLogger(t)
	mockMailer := mailermocks.NewMockMailer(t)

	ts := TestServer{
		Server:          NewServer(mockDBStore, mockErrorLogger, mockInfoLogger, mockMailer),
		MockDBStore:     mockDBStore,
		MockErrorLogger: mockErrorLogger,
		MockInfoLogger:  mockInfoLogger,
		MockMailer:      mockMailer,
	}

	// load web page templates cache
	err := ts.Renderer.LoadGoHtmlPageTemplates()
	require.NoError(t, err)

	// load mail templates cache
	err = ts.Renderer.LoadGoHtmlMailTemplates()
	require.NoError(t, err)

	return &ts
}

// BuildLogAnyErrorStub builds the MockLogger Log() stub for testing of any error logging
func (ts *TestServer) BuildLogAnyErrorStub() {
	ts.MockErrorLogger.On("MyLogChannel").Return(nil).Once()
	ts.MockErrorLogger.On("Log", mock.Anything).Once()
}

// BuildLogErrorStub builds the MockLogger Log() stub for testing of specific error logging
func (ts *TestServer) BuildLogErrorStub(err error) {
	ts.MockErrorLogger.On("MyLogChannel").Return(nil).Once()
	ts.MockErrorLogger.On("Log", err).Once()
}

// BuildLogAnyInfoStub builds the MockLogger Log() stub for testing of any info logging
func (ts *TestServer) BuildLogAnyInfoStub() {
	ts.MockInfoLogger.On("MyLogChannel").Return(nil).Once()
	ts.MockInfoLogger.On("Log", mock.Anything).Once()
}

// BuildLogAnyInfoStub builds the MockLogger Log() stub for testing of specific info logging
func (ts *TestServer) BuildLogInfoStub(info string) {
	ts.MockInfoLogger.On("MyLogChannel").Return(nil).Once()
	ts.MockInfoLogger.On("Log", info).Once()
}

// BuildSendAnyMailStub builds the MockMailer SendMail() stub for testing of any mail sending
func (ts *TestServer) BuildSendAnyMailStub() {
	ts.MockMailer.On("MyMailChannel").Return(nil).Once()
	ts.MockMailer.On("SendMail", mock.Anything).Return(nil).Once()
}

// BuildSendMailStub builds the MockMailer SendMail() stub for testing of specific mail sending
func (ts *TestServer) BuildSendMailStub(data mailers.MailData) {
	ts.MockMailer.On("MyMailChannel").Return(nil).Once()
	ts.MockMailer.On("SendMail", data).Return(nil).Once()
}

// NewTestRequest creates a new get request for use in testing
func (ts *TestServer) NewRequest(method string, url string, body io.Reader) *http.Request {
	return httptest.NewRequest(method, url, body)
}

// NewTestRequestWithSession creates a new get request with new session data for use in testing
func (ts *TestServer) NewRequestWithSession(t *testing.T, method string, url string, body io.Reader) *http.Request {
	// checks that the session manager is loaded
	require.NotNil(t, app.Session)

	// creating new request
	r := httptest.NewRequest(method, url, body)

	if method == http.MethodPost {
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	// adding new session data to context
	ctx, err := app.Session.Load(r.Context(), "X-Session")
	require.NoError(t, err)
	require.NotNil(t, ctx)

	return r.WithContext(ctx)
}

// ServeRequest execute a ServerHTTP method and return the response recorder
func (ts *TestServer) ServeRequest(r *http.Request) *httptest.ResponseRecorder {
	// build request logging stub
	ts.BuildLogAnyInfoStub()

	rr := httptest.NewRecorder()
	ts.Router.Handler.ServeHTTP(rr, r)

	return rr
}
