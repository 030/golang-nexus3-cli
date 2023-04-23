// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewSearchParams creates a new SearchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchParams() *SearchParams {
	return &SearchParams{
		requestTimeout: cr.DefaultTimeout,
	}
}

// NewSearchParamsWithTimeout creates a new SearchParams object
// with the ability to set a timeout on a request.
func NewSearchParamsWithTimeout(timeout time.Duration) *SearchParams {
	return &SearchParams{
		requestTimeout: timeout,
	}
}

// NewSearchParamsWithContext creates a new SearchParams object
// with the ability to set a context for a request.
func NewSearchParamsWithContext(ctx context.Context) *SearchParams {
	return &SearchParams{
		Context: ctx,
	}
}

// NewSearchParamsWithHTTPClient creates a new SearchParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchParamsWithHTTPClient(client *http.Client) *SearchParams {
	return &SearchParams{
		HTTPClient: client,
	}
}

/*
SearchParams contains all the parameters to send to the API endpoint

	for the search operation.

	Typically these are written to a http.Request.
*/
type SearchParams struct {

	/* ConanBaseVersion.

	   Conan base version
	*/
	ConanBaseVersion *string

	/* ConanChannel.

	   Conan channel
	*/
	ConanChannel *string

	/* ConanRevision.

	   Conan recipe revision
	*/
	ConanRevision *string

	/* ContinuationToken.

	   A token returned by a prior request. If present, the next page of results are returned
	*/
	ContinuationToken string

	/* Direction.

	   The direction to sort records in, defaults to ascending ('asc') for all sort fields, except version, which defaults to descending ('desc')
	*/
	Direction string

	/* DockerContentDigest.

	   Docker content digest
	*/
	DockerContentDigest *string

	/* DockerImageName.

	   Docker image name
	*/
	DockerImageName *string

	/* DockerImageTag.

	   Docker image tag
	*/
	DockerImageTag *string

	/* DockerLayerID.

	   Docker layer ID
	*/
	DockerLayerID *string

	/* Format.

	   Query by format
	*/
	Format *string

	/* Gavec.

	   Group asset version extension classifier
	*/
	Gavec *string

	/* Group.

	   Component group
	*/
	Group *string

	/* MavenArtifactID.

	   Maven artifactId
	*/
	MavenArtifactID *string

	/* MavenBaseVersion.

	   Maven base version
	*/
	MavenBaseVersion *string

	/* MavenClassifier.

	   Maven classifier of component's asset
	*/
	MavenClassifier *string

	/* MavenExtension.

	   Maven extension of component's asset
	*/
	MavenExtension *string

	/* MavenGroupID.

	   Maven groupId
	*/
	MavenGroupID *string

	/* Md5.

	   Specific MD5 hash of component's asset
	*/
	Md5 *string

	/* Name.

	   Component name
	*/
	Name *string

	/* NpmAuthor.

	   npm author
	*/
	NpmAuthor *string

	/* NpmDescription.

	   npm description
	*/
	NpmDescription *string

	/* NpmKeywords.

	   npm keywords
	*/
	NpmKeywords *string

	/* NpmLicense.

	   npm license
	*/
	NpmLicense *string

	/* NpmScope.

	   npm scope
	*/
	NpmScope *string

	/* NpmTaggedIs.

	   npm tagged is
	*/
	NpmTaggedIs *string

	/* NpmTaggedNot.

	   npm tagged not
	*/
	NpmTaggedNot *string

	/* NugetAuthors.

	   NuGet authors
	*/
	NugetAuthors *string

	/* NugetDescription.

	   NuGet description
	*/
	NugetDescription *string

	/* NugetID.

	   NuGet id
	*/
	NugetID *string

	/* NugetSummary.

	   NuGet summary
	*/
	NugetSummary *string

	/* NugetTags.

	   NuGet tags
	*/
	NugetTags *string

	/* NugetTitle.

	   NuGet title
	*/
	NugetTitle *string

	/* P2PluginName.

	   p2 plugin name
	*/
	P2PluginName *string

	/* Prerelease.

	   Prerelease version flag
	*/
	Prerelease *string

	/* PypiClassifiers.

	   PyPI classifiers
	*/
	PypiClassifiers *string

	/* PypiDescription.

	   PyPI description
	*/
	PypiDescription *string

	/* PypiKeywords.

	   PyPI keywords
	*/
	PypiKeywords *string

	/* PypiSummary.

	   PyPI summary
	*/
	PypiSummary *string

	/* Q.

	   Query by keyword
	*/
	Q *string

	/* Repository.

	   Repository name
	*/
	Repository *string

	/* RubygemsDescription.

	   RubyGems description
	*/
	RubygemsDescription *string

	/* RubygemsPlatform.

	   RubyGems platform
	*/
	RubygemsPlatform *string

	/* RubygemsSummary.

	   RubyGems summary
	*/
	RubygemsSummary *string

	/* Sha1.

	   Specific SHA-1 hash of component's asset
	*/
	Sha1 *string

	/* Sha256.

	   Specific SHA-256 hash of component's asset
	*/
	Sha256 *string

	/* Sha512.

	   Specific SHA-512 hash of component's asset
	*/
	Sha512 *string

	/* Sort.

	   The field to sort the results against, if left empty, a sort based on match weight will be used.
	*/
	Sort string

	/* Timeout.

	   How long to wait for search results in seconds. If this value is not provided, the system default timeout will be used.

	   Format: int32
	*/
	Timeout int32

	/* Version.

	   Component version
	*/
	Version *string

	/* YumArchitecture.

	   Yum architecture
	*/
	YumArchitecture *string

	/* YumName.

	   Yum package name
	*/
	YumName *string

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithDefaults hydrates default values in the search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchParams) WithDefaults() *SearchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithRequestTimeout adds the timeout to the search params
func (o *SearchParams) WithRequestTimeout(timeout time.Duration) *SearchParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the search params
func (o *SearchParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the search params
func (o *SearchParams) WithContext(ctx context.Context) *SearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search params
func (o *SearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search params
func (o *SearchParams) WithHTTPClient(client *http.Client) *SearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search params
func (o *SearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConanBaseVersion adds the conanBaseVersion to the search params
func (o *SearchParams) WithConanBaseVersion(conanBaseVersion *string) *SearchParams {
	o.SetConanBaseVersion(conanBaseVersion)
	return o
}

// SetConanBaseVersion adds the conanBaseVersion to the search params
func (o *SearchParams) SetConanBaseVersion(conanBaseVersion *string) {
	o.ConanBaseVersion = conanBaseVersion
}

// WithConanChannel adds the conanChannel to the search params
func (o *SearchParams) WithConanChannel(conanChannel *string) *SearchParams {
	o.SetConanChannel(conanChannel)
	return o
}

// SetConanChannel adds the conanChannel to the search params
func (o *SearchParams) SetConanChannel(conanChannel *string) {
	o.ConanChannel = conanChannel
}

// WithConanRevision adds the conanRevision to the search params
func (o *SearchParams) WithConanRevision(conanRevision *string) *SearchParams {
	o.SetConanRevision(conanRevision)
	return o
}

// SetConanRevision adds the conanRevision to the search params
func (o *SearchParams) SetConanRevision(conanRevision *string) {
	o.ConanRevision = conanRevision
}

// WithContinuationToken adds the continuationToken to the search params
func (o *SearchParams) WithContinuationToken(continuationToken string) *SearchParams {
	o.SetContinuationToken(continuationToken)
	return o
}

// SetContinuationToken adds the continuationToken to the search params
func (o *SearchParams) SetContinuationToken(continuationToken string) {
	o.ContinuationToken = continuationToken
}

// WithDirection adds the direction to the search params
func (o *SearchParams) WithDirection(direction string) *SearchParams {
	o.SetDirection(direction)
	return o
}

// SetDirection adds the direction to the search params
func (o *SearchParams) SetDirection(direction string) {
	o.Direction = direction
}

// WithDockerContentDigest adds the dockerContentDigest to the search params
func (o *SearchParams) WithDockerContentDigest(dockerContentDigest *string) *SearchParams {
	o.SetDockerContentDigest(dockerContentDigest)
	return o
}

// SetDockerContentDigest adds the dockerContentDigest to the search params
func (o *SearchParams) SetDockerContentDigest(dockerContentDigest *string) {
	o.DockerContentDigest = dockerContentDigest
}

// WithDockerImageName adds the dockerImageName to the search params
func (o *SearchParams) WithDockerImageName(dockerImageName *string) *SearchParams {
	o.SetDockerImageName(dockerImageName)
	return o
}

// SetDockerImageName adds the dockerImageName to the search params
func (o *SearchParams) SetDockerImageName(dockerImageName *string) {
	o.DockerImageName = dockerImageName
}

// WithDockerImageTag adds the dockerImageTag to the search params
func (o *SearchParams) WithDockerImageTag(dockerImageTag *string) *SearchParams {
	o.SetDockerImageTag(dockerImageTag)
	return o
}

// SetDockerImageTag adds the dockerImageTag to the search params
func (o *SearchParams) SetDockerImageTag(dockerImageTag *string) {
	o.DockerImageTag = dockerImageTag
}

// WithDockerLayerID adds the dockerLayerID to the search params
func (o *SearchParams) WithDockerLayerID(dockerLayerID *string) *SearchParams {
	o.SetDockerLayerID(dockerLayerID)
	return o
}

// SetDockerLayerID adds the dockerLayerId to the search params
func (o *SearchParams) SetDockerLayerID(dockerLayerID *string) {
	o.DockerLayerID = dockerLayerID
}

// WithFormat adds the format to the search params
func (o *SearchParams) WithFormat(format *string) *SearchParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the search params
func (o *SearchParams) SetFormat(format *string) {
	o.Format = format
}

// WithGavec adds the gavec to the search params
func (o *SearchParams) WithGavec(gavec *string) *SearchParams {
	o.SetGavec(gavec)
	return o
}

// SetGavec adds the gavec to the search params
func (o *SearchParams) SetGavec(gavec *string) {
	o.Gavec = gavec
}

// WithGroup adds the group to the search params
func (o *SearchParams) WithGroup(group *string) *SearchParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the search params
func (o *SearchParams) SetGroup(group *string) {
	o.Group = group
}

// WithMavenArtifactID adds the mavenArtifactID to the search params
func (o *SearchParams) WithMavenArtifactID(mavenArtifactID *string) *SearchParams {
	o.SetMavenArtifactID(mavenArtifactID)
	return o
}

// SetMavenArtifactID adds the mavenArtifactId to the search params
func (o *SearchParams) SetMavenArtifactID(mavenArtifactID *string) {
	o.MavenArtifactID = mavenArtifactID
}

// WithMavenBaseVersion adds the mavenBaseVersion to the search params
func (o *SearchParams) WithMavenBaseVersion(mavenBaseVersion *string) *SearchParams {
	o.SetMavenBaseVersion(mavenBaseVersion)
	return o
}

// SetMavenBaseVersion adds the mavenBaseVersion to the search params
func (o *SearchParams) SetMavenBaseVersion(mavenBaseVersion *string) {
	o.MavenBaseVersion = mavenBaseVersion
}

// WithMavenClassifier adds the mavenClassifier to the search params
func (o *SearchParams) WithMavenClassifier(mavenClassifier *string) *SearchParams {
	o.SetMavenClassifier(mavenClassifier)
	return o
}

// SetMavenClassifier adds the mavenClassifier to the search params
func (o *SearchParams) SetMavenClassifier(mavenClassifier *string) {
	o.MavenClassifier = mavenClassifier
}

// WithMavenExtension adds the mavenExtension to the search params
func (o *SearchParams) WithMavenExtension(mavenExtension *string) *SearchParams {
	o.SetMavenExtension(mavenExtension)
	return o
}

// SetMavenExtension adds the mavenExtension to the search params
func (o *SearchParams) SetMavenExtension(mavenExtension *string) {
	o.MavenExtension = mavenExtension
}

// WithMavenGroupID adds the mavenGroupID to the search params
func (o *SearchParams) WithMavenGroupID(mavenGroupID *string) *SearchParams {
	o.SetMavenGroupID(mavenGroupID)
	return o
}

// SetMavenGroupID adds the mavenGroupId to the search params
func (o *SearchParams) SetMavenGroupID(mavenGroupID *string) {
	o.MavenGroupID = mavenGroupID
}

// WithMd5 adds the md5 to the search params
func (o *SearchParams) WithMd5(md5 *string) *SearchParams {
	o.SetMd5(md5)
	return o
}

// SetMd5 adds the md5 to the search params
func (o *SearchParams) SetMd5(md5 *string) {
	o.Md5 = md5
}

// WithName adds the name to the search params
func (o *SearchParams) WithName(name *string) *SearchParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the search params
func (o *SearchParams) SetName(name *string) {
	o.Name = name
}

// WithNpmAuthor adds the npmAuthor to the search params
func (o *SearchParams) WithNpmAuthor(npmAuthor *string) *SearchParams {
	o.SetNpmAuthor(npmAuthor)
	return o
}

// SetNpmAuthor adds the npmAuthor to the search params
func (o *SearchParams) SetNpmAuthor(npmAuthor *string) {
	o.NpmAuthor = npmAuthor
}

// WithNpmDescription adds the npmDescription to the search params
func (o *SearchParams) WithNpmDescription(npmDescription *string) *SearchParams {
	o.SetNpmDescription(npmDescription)
	return o
}

// SetNpmDescription adds the npmDescription to the search params
func (o *SearchParams) SetNpmDescription(npmDescription *string) {
	o.NpmDescription = npmDescription
}

// WithNpmKeywords adds the npmKeywords to the search params
func (o *SearchParams) WithNpmKeywords(npmKeywords *string) *SearchParams {
	o.SetNpmKeywords(npmKeywords)
	return o
}

// SetNpmKeywords adds the npmKeywords to the search params
func (o *SearchParams) SetNpmKeywords(npmKeywords *string) {
	o.NpmKeywords = npmKeywords
}

// WithNpmLicense adds the npmLicense to the search params
func (o *SearchParams) WithNpmLicense(npmLicense *string) *SearchParams {
	o.SetNpmLicense(npmLicense)
	return o
}

// SetNpmLicense adds the npmLicense to the search params
func (o *SearchParams) SetNpmLicense(npmLicense *string) {
	o.NpmLicense = npmLicense
}

// WithNpmScope adds the npmScope to the search params
func (o *SearchParams) WithNpmScope(npmScope *string) *SearchParams {
	o.SetNpmScope(npmScope)
	return o
}

// SetNpmScope adds the npmScope to the search params
func (o *SearchParams) SetNpmScope(npmScope *string) {
	o.NpmScope = npmScope
}

// WithNpmTaggedIs adds the npmTaggedIs to the search params
func (o *SearchParams) WithNpmTaggedIs(npmTaggedIs *string) *SearchParams {
	o.SetNpmTaggedIs(npmTaggedIs)
	return o
}

// SetNpmTaggedIs adds the npmTaggedIs to the search params
func (o *SearchParams) SetNpmTaggedIs(npmTaggedIs *string) {
	o.NpmTaggedIs = npmTaggedIs
}

// WithNpmTaggedNot adds the npmTaggedNot to the search params
func (o *SearchParams) WithNpmTaggedNot(npmTaggedNot *string) *SearchParams {
	o.SetNpmTaggedNot(npmTaggedNot)
	return o
}

// SetNpmTaggedNot adds the npmTaggedNot to the search params
func (o *SearchParams) SetNpmTaggedNot(npmTaggedNot *string) {
	o.NpmTaggedNot = npmTaggedNot
}

// WithNugetAuthors adds the nugetAuthors to the search params
func (o *SearchParams) WithNugetAuthors(nugetAuthors *string) *SearchParams {
	o.SetNugetAuthors(nugetAuthors)
	return o
}

// SetNugetAuthors adds the nugetAuthors to the search params
func (o *SearchParams) SetNugetAuthors(nugetAuthors *string) {
	o.NugetAuthors = nugetAuthors
}

// WithNugetDescription adds the nugetDescription to the search params
func (o *SearchParams) WithNugetDescription(nugetDescription *string) *SearchParams {
	o.SetNugetDescription(nugetDescription)
	return o
}

// SetNugetDescription adds the nugetDescription to the search params
func (o *SearchParams) SetNugetDescription(nugetDescription *string) {
	o.NugetDescription = nugetDescription
}

// WithNugetID adds the nugetID to the search params
func (o *SearchParams) WithNugetID(nugetID *string) *SearchParams {
	o.SetNugetID(nugetID)
	return o
}

// SetNugetID adds the nugetId to the search params
func (o *SearchParams) SetNugetID(nugetID *string) {
	o.NugetID = nugetID
}

// WithNugetSummary adds the nugetSummary to the search params
func (o *SearchParams) WithNugetSummary(nugetSummary *string) *SearchParams {
	o.SetNugetSummary(nugetSummary)
	return o
}

// SetNugetSummary adds the nugetSummary to the search params
func (o *SearchParams) SetNugetSummary(nugetSummary *string) {
	o.NugetSummary = nugetSummary
}

// WithNugetTags adds the nugetTags to the search params
func (o *SearchParams) WithNugetTags(nugetTags *string) *SearchParams {
	o.SetNugetTags(nugetTags)
	return o
}

// SetNugetTags adds the nugetTags to the search params
func (o *SearchParams) SetNugetTags(nugetTags *string) {
	o.NugetTags = nugetTags
}

// WithNugetTitle adds the nugetTitle to the search params
func (o *SearchParams) WithNugetTitle(nugetTitle *string) *SearchParams {
	o.SetNugetTitle(nugetTitle)
	return o
}

// SetNugetTitle adds the nugetTitle to the search params
func (o *SearchParams) SetNugetTitle(nugetTitle *string) {
	o.NugetTitle = nugetTitle
}

// WithP2PluginName adds the p2PluginName to the search params
func (o *SearchParams) WithP2PluginName(p2PluginName *string) *SearchParams {
	o.SetP2PluginName(p2PluginName)
	return o
}

// SetP2PluginName adds the p2PluginName to the search params
func (o *SearchParams) SetP2PluginName(p2PluginName *string) {
	o.P2PluginName = p2PluginName
}

// WithPrerelease adds the prerelease to the search params
func (o *SearchParams) WithPrerelease(prerelease *string) *SearchParams {
	o.SetPrerelease(prerelease)
	return o
}

// SetPrerelease adds the prerelease to the search params
func (o *SearchParams) SetPrerelease(prerelease *string) {
	o.Prerelease = prerelease
}

// WithPypiClassifiers adds the pypiClassifiers to the search params
func (o *SearchParams) WithPypiClassifiers(pypiClassifiers *string) *SearchParams {
	o.SetPypiClassifiers(pypiClassifiers)
	return o
}

// SetPypiClassifiers adds the pypiClassifiers to the search params
func (o *SearchParams) SetPypiClassifiers(pypiClassifiers *string) {
	o.PypiClassifiers = pypiClassifiers
}

// WithPypiDescription adds the pypiDescription to the search params
func (o *SearchParams) WithPypiDescription(pypiDescription *string) *SearchParams {
	o.SetPypiDescription(pypiDescription)
	return o
}

// SetPypiDescription adds the pypiDescription to the search params
func (o *SearchParams) SetPypiDescription(pypiDescription *string) {
	o.PypiDescription = pypiDescription
}

// WithPypiKeywords adds the pypiKeywords to the search params
func (o *SearchParams) WithPypiKeywords(pypiKeywords *string) *SearchParams {
	o.SetPypiKeywords(pypiKeywords)
	return o
}

// SetPypiKeywords adds the pypiKeywords to the search params
func (o *SearchParams) SetPypiKeywords(pypiKeywords *string) {
	o.PypiKeywords = pypiKeywords
}

// WithPypiSummary adds the pypiSummary to the search params
func (o *SearchParams) WithPypiSummary(pypiSummary *string) *SearchParams {
	o.SetPypiSummary(pypiSummary)
	return o
}

// SetPypiSummary adds the pypiSummary to the search params
func (o *SearchParams) SetPypiSummary(pypiSummary *string) {
	o.PypiSummary = pypiSummary
}

// WithQ adds the q to the search params
func (o *SearchParams) WithQ(q *string) *SearchParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the search params
func (o *SearchParams) SetQ(q *string) {
	o.Q = q
}

// WithRepository adds the repository to the search params
func (o *SearchParams) WithRepository(repository *string) *SearchParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the search params
func (o *SearchParams) SetRepository(repository *string) {
	o.Repository = repository
}

// WithRubygemsDescription adds the rubygemsDescription to the search params
func (o *SearchParams) WithRubygemsDescription(rubygemsDescription *string) *SearchParams {
	o.SetRubygemsDescription(rubygemsDescription)
	return o
}

// SetRubygemsDescription adds the rubygemsDescription to the search params
func (o *SearchParams) SetRubygemsDescription(rubygemsDescription *string) {
	o.RubygemsDescription = rubygemsDescription
}

// WithRubygemsPlatform adds the rubygemsPlatform to the search params
func (o *SearchParams) WithRubygemsPlatform(rubygemsPlatform *string) *SearchParams {
	o.SetRubygemsPlatform(rubygemsPlatform)
	return o
}

// SetRubygemsPlatform adds the rubygemsPlatform to the search params
func (o *SearchParams) SetRubygemsPlatform(rubygemsPlatform *string) {
	o.RubygemsPlatform = rubygemsPlatform
}

// WithRubygemsSummary adds the rubygemsSummary to the search params
func (o *SearchParams) WithRubygemsSummary(rubygemsSummary *string) *SearchParams {
	o.SetRubygemsSummary(rubygemsSummary)
	return o
}

// SetRubygemsSummary adds the rubygemsSummary to the search params
func (o *SearchParams) SetRubygemsSummary(rubygemsSummary *string) {
	o.RubygemsSummary = rubygemsSummary
}

// WithSha1 adds the sha1 to the search params
func (o *SearchParams) WithSha1(sha1 *string) *SearchParams {
	o.SetSha1(sha1)
	return o
}

// SetSha1 adds the sha1 to the search params
func (o *SearchParams) SetSha1(sha1 *string) {
	o.Sha1 = sha1
}

// WithSha256 adds the sha256 to the search params
func (o *SearchParams) WithSha256(sha256 *string) *SearchParams {
	o.SetSha256(sha256)
	return o
}

// SetSha256 adds the sha256 to the search params
func (o *SearchParams) SetSha256(sha256 *string) {
	o.Sha256 = sha256
}

// WithSha512 adds the sha512 to the search params
func (o *SearchParams) WithSha512(sha512 *string) *SearchParams {
	o.SetSha512(sha512)
	return o
}

// SetSha512 adds the sha512 to the search params
func (o *SearchParams) SetSha512(sha512 *string) {
	o.Sha512 = sha512
}

// WithSort adds the sort to the search params
func (o *SearchParams) WithSort(sort string) *SearchParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the search params
func (o *SearchParams) SetSort(sort string) {
	o.Sort = sort
}

// WithTimeout adds the timeout to the search params
func (o *SearchParams) WithTimeout(timeout int32) *SearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search params
func (o *SearchParams) SetTimeout(timeout int32) {
	o.Timeout = timeout
}

// WithVersion adds the version to the search params
func (o *SearchParams) WithVersion(version *string) *SearchParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the search params
func (o *SearchParams) SetVersion(version *string) {
	o.Version = version
}

// WithYumArchitecture adds the yumArchitecture to the search params
func (o *SearchParams) WithYumArchitecture(yumArchitecture *string) *SearchParams {
	o.SetYumArchitecture(yumArchitecture)
	return o
}

// SetYumArchitecture adds the yumArchitecture to the search params
func (o *SearchParams) SetYumArchitecture(yumArchitecture *string) {
	o.YumArchitecture = yumArchitecture
}

// WithYumName adds the yumName to the search params
func (o *SearchParams) WithYumName(yumName *string) *SearchParams {
	o.SetYumName(yumName)
	return o
}

// SetYumName adds the yumName to the search params
func (o *SearchParams) SetYumName(yumName *string) {
	o.YumName = yumName
}

// WriteToRequest writes these params to a swagger request
func (o *SearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.requestTimeout); err != nil {
		return err
	}
	var res []error

	if o.ConanBaseVersion != nil {

		// query param conan.baseVersion
		var qrConanBaseVersion string

		if o.ConanBaseVersion != nil {
			qrConanBaseVersion = *o.ConanBaseVersion
		}
		qConanBaseVersion := qrConanBaseVersion
		if qConanBaseVersion != "" {

			if err := r.SetQueryParam("conan.baseVersion", qConanBaseVersion); err != nil {
				return err
			}
		}
	}

	if o.ConanChannel != nil {

		// query param conan.channel
		var qrConanChannel string

		if o.ConanChannel != nil {
			qrConanChannel = *o.ConanChannel
		}
		qConanChannel := qrConanChannel
		if qConanChannel != "" {

			if err := r.SetQueryParam("conan.channel", qConanChannel); err != nil {
				return err
			}
		}
	}

	if o.ConanRevision != nil {

		// query param conan.revision
		var qrConanRevision string

		if o.ConanRevision != nil {
			qrConanRevision = *o.ConanRevision
		}
		qConanRevision := qrConanRevision
		if qConanRevision != "" {

			if err := r.SetQueryParam("conan.revision", qConanRevision); err != nil {
				return err
			}
		}
	}

	// query param continuationToken
	qrContinuationToken := o.ContinuationToken
	qContinuationToken := qrContinuationToken

	if err := r.SetQueryParam("continuationToken", qContinuationToken); err != nil {
		return err
	}

	// query param direction
	qrDirection := o.Direction
	qDirection := qrDirection

	if err := r.SetQueryParam("direction", qDirection); err != nil {
		return err
	}

	if o.DockerContentDigest != nil {

		// query param docker.contentDigest
		var qrDockerContentDigest string

		if o.DockerContentDigest != nil {
			qrDockerContentDigest = *o.DockerContentDigest
		}
		qDockerContentDigest := qrDockerContentDigest
		if qDockerContentDigest != "" {

			if err := r.SetQueryParam("docker.contentDigest", qDockerContentDigest); err != nil {
				return err
			}
		}
	}

	if o.DockerImageName != nil {

		// query param docker.imageName
		var qrDockerImageName string

		if o.DockerImageName != nil {
			qrDockerImageName = *o.DockerImageName
		}
		qDockerImageName := qrDockerImageName
		if qDockerImageName != "" {

			if err := r.SetQueryParam("docker.imageName", qDockerImageName); err != nil {
				return err
			}
		}
	}

	if o.DockerImageTag != nil {

		// query param docker.imageTag
		var qrDockerImageTag string

		if o.DockerImageTag != nil {
			qrDockerImageTag = *o.DockerImageTag
		}
		qDockerImageTag := qrDockerImageTag
		if qDockerImageTag != "" {

			if err := r.SetQueryParam("docker.imageTag", qDockerImageTag); err != nil {
				return err
			}
		}
	}

	if o.DockerLayerID != nil {

		// query param docker.layerId
		var qrDockerLayerID string

		if o.DockerLayerID != nil {
			qrDockerLayerID = *o.DockerLayerID
		}
		qDockerLayerID := qrDockerLayerID
		if qDockerLayerID != "" {

			if err := r.SetQueryParam("docker.layerId", qDockerLayerID); err != nil {
				return err
			}
		}
	}

	if o.Format != nil {

		// query param format
		var qrFormat string

		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {

			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}
	}

	if o.Gavec != nil {

		// query param gavec
		var qrGavec string

		if o.Gavec != nil {
			qrGavec = *o.Gavec
		}
		qGavec := qrGavec
		if qGavec != "" {

			if err := r.SetQueryParam("gavec", qGavec); err != nil {
				return err
			}
		}
	}

	if o.Group != nil {

		// query param group
		var qrGroup string

		if o.Group != nil {
			qrGroup = *o.Group
		}
		qGroup := qrGroup
		if qGroup != "" {

			if err := r.SetQueryParam("group", qGroup); err != nil {
				return err
			}
		}
	}

	if o.MavenArtifactID != nil {

		// query param maven.artifactId
		var qrMavenArtifactID string

		if o.MavenArtifactID != nil {
			qrMavenArtifactID = *o.MavenArtifactID
		}
		qMavenArtifactID := qrMavenArtifactID
		if qMavenArtifactID != "" {

			if err := r.SetQueryParam("maven.artifactId", qMavenArtifactID); err != nil {
				return err
			}
		}
	}

	if o.MavenBaseVersion != nil {

		// query param maven.baseVersion
		var qrMavenBaseVersion string

		if o.MavenBaseVersion != nil {
			qrMavenBaseVersion = *o.MavenBaseVersion
		}
		qMavenBaseVersion := qrMavenBaseVersion
		if qMavenBaseVersion != "" {

			if err := r.SetQueryParam("maven.baseVersion", qMavenBaseVersion); err != nil {
				return err
			}
		}
	}

	if o.MavenClassifier != nil {

		// query param maven.classifier
		var qrMavenClassifier string

		if o.MavenClassifier != nil {
			qrMavenClassifier = *o.MavenClassifier
		}
		qMavenClassifier := qrMavenClassifier
		if qMavenClassifier != "" {

			if err := r.SetQueryParam("maven.classifier", qMavenClassifier); err != nil {
				return err
			}
		}
	}

	if o.MavenExtension != nil {

		// query param maven.extension
		var qrMavenExtension string

		if o.MavenExtension != nil {
			qrMavenExtension = *o.MavenExtension
		}
		qMavenExtension := qrMavenExtension
		if qMavenExtension != "" {

			if err := r.SetQueryParam("maven.extension", qMavenExtension); err != nil {
				return err
			}
		}
	}

	if o.MavenGroupID != nil {

		// query param maven.groupId
		var qrMavenGroupID string

		if o.MavenGroupID != nil {
			qrMavenGroupID = *o.MavenGroupID
		}
		qMavenGroupID := qrMavenGroupID
		if qMavenGroupID != "" {

			if err := r.SetQueryParam("maven.groupId", qMavenGroupID); err != nil {
				return err
			}
		}
	}

	if o.Md5 != nil {

		// query param md5
		var qrMd5 string

		if o.Md5 != nil {
			qrMd5 = *o.Md5
		}
		qMd5 := qrMd5
		if qMd5 != "" {

			if err := r.SetQueryParam("md5", qMd5); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.NpmAuthor != nil {

		// query param npm.author
		var qrNpmAuthor string

		if o.NpmAuthor != nil {
			qrNpmAuthor = *o.NpmAuthor
		}
		qNpmAuthor := qrNpmAuthor
		if qNpmAuthor != "" {

			if err := r.SetQueryParam("npm.author", qNpmAuthor); err != nil {
				return err
			}
		}
	}

	if o.NpmDescription != nil {

		// query param npm.description
		var qrNpmDescription string

		if o.NpmDescription != nil {
			qrNpmDescription = *o.NpmDescription
		}
		qNpmDescription := qrNpmDescription
		if qNpmDescription != "" {

			if err := r.SetQueryParam("npm.description", qNpmDescription); err != nil {
				return err
			}
		}
	}

	if o.NpmKeywords != nil {

		// query param npm.keywords
		var qrNpmKeywords string

		if o.NpmKeywords != nil {
			qrNpmKeywords = *o.NpmKeywords
		}
		qNpmKeywords := qrNpmKeywords
		if qNpmKeywords != "" {

			if err := r.SetQueryParam("npm.keywords", qNpmKeywords); err != nil {
				return err
			}
		}
	}

	if o.NpmLicense != nil {

		// query param npm.license
		var qrNpmLicense string

		if o.NpmLicense != nil {
			qrNpmLicense = *o.NpmLicense
		}
		qNpmLicense := qrNpmLicense
		if qNpmLicense != "" {

			if err := r.SetQueryParam("npm.license", qNpmLicense); err != nil {
				return err
			}
		}
	}

	if o.NpmScope != nil {

		// query param npm.scope
		var qrNpmScope string

		if o.NpmScope != nil {
			qrNpmScope = *o.NpmScope
		}
		qNpmScope := qrNpmScope
		if qNpmScope != "" {

			if err := r.SetQueryParam("npm.scope", qNpmScope); err != nil {
				return err
			}
		}
	}

	if o.NpmTaggedIs != nil {

		// query param npm.tagged_is
		var qrNpmTaggedIs string

		if o.NpmTaggedIs != nil {
			qrNpmTaggedIs = *o.NpmTaggedIs
		}
		qNpmTaggedIs := qrNpmTaggedIs
		if qNpmTaggedIs != "" {

			if err := r.SetQueryParam("npm.tagged_is", qNpmTaggedIs); err != nil {
				return err
			}
		}
	}

	if o.NpmTaggedNot != nil {

		// query param npm.tagged_not
		var qrNpmTaggedNot string

		if o.NpmTaggedNot != nil {
			qrNpmTaggedNot = *o.NpmTaggedNot
		}
		qNpmTaggedNot := qrNpmTaggedNot
		if qNpmTaggedNot != "" {

			if err := r.SetQueryParam("npm.tagged_not", qNpmTaggedNot); err != nil {
				return err
			}
		}
	}

	if o.NugetAuthors != nil {

		// query param nuget.authors
		var qrNugetAuthors string

		if o.NugetAuthors != nil {
			qrNugetAuthors = *o.NugetAuthors
		}
		qNugetAuthors := qrNugetAuthors
		if qNugetAuthors != "" {

			if err := r.SetQueryParam("nuget.authors", qNugetAuthors); err != nil {
				return err
			}
		}
	}

	if o.NugetDescription != nil {

		// query param nuget.description
		var qrNugetDescription string

		if o.NugetDescription != nil {
			qrNugetDescription = *o.NugetDescription
		}
		qNugetDescription := qrNugetDescription
		if qNugetDescription != "" {

			if err := r.SetQueryParam("nuget.description", qNugetDescription); err != nil {
				return err
			}
		}
	}

	if o.NugetID != nil {

		// query param nuget.id
		var qrNugetID string

		if o.NugetID != nil {
			qrNugetID = *o.NugetID
		}
		qNugetID := qrNugetID
		if qNugetID != "" {

			if err := r.SetQueryParam("nuget.id", qNugetID); err != nil {
				return err
			}
		}
	}

	if o.NugetSummary != nil {

		// query param nuget.summary
		var qrNugetSummary string

		if o.NugetSummary != nil {
			qrNugetSummary = *o.NugetSummary
		}
		qNugetSummary := qrNugetSummary
		if qNugetSummary != "" {

			if err := r.SetQueryParam("nuget.summary", qNugetSummary); err != nil {
				return err
			}
		}
	}

	if o.NugetTags != nil {

		// query param nuget.tags
		var qrNugetTags string

		if o.NugetTags != nil {
			qrNugetTags = *o.NugetTags
		}
		qNugetTags := qrNugetTags
		if qNugetTags != "" {

			if err := r.SetQueryParam("nuget.tags", qNugetTags); err != nil {
				return err
			}
		}
	}

	if o.NugetTitle != nil {

		// query param nuget.title
		var qrNugetTitle string

		if o.NugetTitle != nil {
			qrNugetTitle = *o.NugetTitle
		}
		qNugetTitle := qrNugetTitle
		if qNugetTitle != "" {

			if err := r.SetQueryParam("nuget.title", qNugetTitle); err != nil {
				return err
			}
		}
	}

	if o.P2PluginName != nil {

		// query param p2.pluginName
		var qrP2PluginName string

		if o.P2PluginName != nil {
			qrP2PluginName = *o.P2PluginName
		}
		qP2PluginName := qrP2PluginName
		if qP2PluginName != "" {

			if err := r.SetQueryParam("p2.pluginName", qP2PluginName); err != nil {
				return err
			}
		}
	}

	if o.Prerelease != nil {

		// query param prerelease
		var qrPrerelease string

		if o.Prerelease != nil {
			qrPrerelease = *o.Prerelease
		}
		qPrerelease := qrPrerelease
		if qPrerelease != "" {

			if err := r.SetQueryParam("prerelease", qPrerelease); err != nil {
				return err
			}
		}
	}

	if o.PypiClassifiers != nil {

		// query param pypi.classifiers
		var qrPypiClassifiers string

		if o.PypiClassifiers != nil {
			qrPypiClassifiers = *o.PypiClassifiers
		}
		qPypiClassifiers := qrPypiClassifiers
		if qPypiClassifiers != "" {

			if err := r.SetQueryParam("pypi.classifiers", qPypiClassifiers); err != nil {
				return err
			}
		}
	}

	if o.PypiDescription != nil {

		// query param pypi.description
		var qrPypiDescription string

		if o.PypiDescription != nil {
			qrPypiDescription = *o.PypiDescription
		}
		qPypiDescription := qrPypiDescription
		if qPypiDescription != "" {

			if err := r.SetQueryParam("pypi.description", qPypiDescription); err != nil {
				return err
			}
		}
	}

	if o.PypiKeywords != nil {

		// query param pypi.keywords
		var qrPypiKeywords string

		if o.PypiKeywords != nil {
			qrPypiKeywords = *o.PypiKeywords
		}
		qPypiKeywords := qrPypiKeywords
		if qPypiKeywords != "" {

			if err := r.SetQueryParam("pypi.keywords", qPypiKeywords); err != nil {
				return err
			}
		}
	}

	if o.PypiSummary != nil {

		// query param pypi.summary
		var qrPypiSummary string

		if o.PypiSummary != nil {
			qrPypiSummary = *o.PypiSummary
		}
		qPypiSummary := qrPypiSummary
		if qPypiSummary != "" {

			if err := r.SetQueryParam("pypi.summary", qPypiSummary); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.Repository != nil {

		// query param repository
		var qrRepository string

		if o.Repository != nil {
			qrRepository = *o.Repository
		}
		qRepository := qrRepository
		if qRepository != "" {

			if err := r.SetQueryParam("repository", qRepository); err != nil {
				return err
			}
		}
	}

	if o.RubygemsDescription != nil {

		// query param rubygems.description
		var qrRubygemsDescription string

		if o.RubygemsDescription != nil {
			qrRubygemsDescription = *o.RubygemsDescription
		}
		qRubygemsDescription := qrRubygemsDescription
		if qRubygemsDescription != "" {

			if err := r.SetQueryParam("rubygems.description", qRubygemsDescription); err != nil {
				return err
			}
		}
	}

	if o.RubygemsPlatform != nil {

		// query param rubygems.platform
		var qrRubygemsPlatform string

		if o.RubygemsPlatform != nil {
			qrRubygemsPlatform = *o.RubygemsPlatform
		}
		qRubygemsPlatform := qrRubygemsPlatform
		if qRubygemsPlatform != "" {

			if err := r.SetQueryParam("rubygems.platform", qRubygemsPlatform); err != nil {
				return err
			}
		}
	}

	if o.RubygemsSummary != nil {

		// query param rubygems.summary
		var qrRubygemsSummary string

		if o.RubygemsSummary != nil {
			qrRubygemsSummary = *o.RubygemsSummary
		}
		qRubygemsSummary := qrRubygemsSummary
		if qRubygemsSummary != "" {

			if err := r.SetQueryParam("rubygems.summary", qRubygemsSummary); err != nil {
				return err
			}
		}
	}

	if o.Sha1 != nil {

		// query param sha1
		var qrSha1 string

		if o.Sha1 != nil {
			qrSha1 = *o.Sha1
		}
		qSha1 := qrSha1
		if qSha1 != "" {

			if err := r.SetQueryParam("sha1", qSha1); err != nil {
				return err
			}
		}
	}

	if o.Sha256 != nil {

		// query param sha256
		var qrSha256 string

		if o.Sha256 != nil {
			qrSha256 = *o.Sha256
		}
		qSha256 := qrSha256
		if qSha256 != "" {

			if err := r.SetQueryParam("sha256", qSha256); err != nil {
				return err
			}
		}
	}

	if o.Sha512 != nil {

		// query param sha512
		var qrSha512 string

		if o.Sha512 != nil {
			qrSha512 = *o.Sha512
		}
		qSha512 := qrSha512
		if qSha512 != "" {

			if err := r.SetQueryParam("sha512", qSha512); err != nil {
				return err
			}
		}
	}

	// query param sort
	qrSort := o.Sort
	qSort := qrSort

	if err := r.SetQueryParam("sort", qSort); err != nil {
		return err
	}

	// query param timeout
	qrTimeout := o.Timeout
	qTimeout := swag.FormatInt32(qrTimeout)

	if err := r.SetQueryParam("timeout", qTimeout); err != nil {
		return err
	}

	if o.Version != nil {

		// query param version
		var qrVersion string

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := qrVersion
		if qVersion != "" {

			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}
	}

	if o.YumArchitecture != nil {

		// query param yum.architecture
		var qrYumArchitecture string

		if o.YumArchitecture != nil {
			qrYumArchitecture = *o.YumArchitecture
		}
		qYumArchitecture := qrYumArchitecture
		if qYumArchitecture != "" {

			if err := r.SetQueryParam("yum.architecture", qYumArchitecture); err != nil {
				return err
			}
		}
	}

	if o.YumName != nil {

		// query param yum.name
		var qrYumName string

		if o.YumName != nil {
			qrYumName = *o.YumName
		}
		qYumName := qrYumName
		if qYumName != "" {

			if err := r.SetQueryParam("yum.name", qYumName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
