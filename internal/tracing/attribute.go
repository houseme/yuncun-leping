/*
 * Copyright yuncun-leping Author(https://houseme.github.io/yuncun-leping/). All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * You can obtain one at https://github.com/houseme/yuncun-leping.
 *
 */

package tracing

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gtrace"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
	"go.opentelemetry.io/otel/trace"
)

const telemetrySDKName = "opentelemetry"

// SetAttributes set tracing attributes
func SetAttributes(r *ghttp.Request, span *gtrace.Span) {
	span.SetAttributes(semconv.HTTPURL(r.URL.Path))
	span.SetAttributes(semconv.HTTPMethod(r.Method))
	span.SetAttributes(semconv.NetHostName(r.GetHost()))
	span.SetAttributes(semconv.HTTPScheme(r.Proto))
	span.SetAttributes(semconv.HTTPStatusCode(r.Response.Status))
	span.SetAttributes(semconv.UserAgentOriginal(r.UserAgent()))
}

// CommonEventOption common event option
func CommonEventOption(_ context.Context, namespace string) trace.SpanStartEventOption {
	return trace.WithAttributes(
		semconv.ServiceNamespace(namespace),
		semconv.TelemetrySDKName(telemetrySDKName),
		semconv.TelemetrySDKVersion("1.0.0"),
		semconv.TelemetryAutoVersion("1.0.0"),
		semconv.TelemetrySDKLanguageGo,
	)
}
