/* 
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * OpenAPI spec version: 0.1
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
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
 */

package swagger

// Derived describes the derived image portion (Occurrence) of the DockerImage relationship.  This image would be produced from a Dockerfile with FROM <DockerImage.Basis in attached Note>.
type Derived struct {

	// The fingerprint of the derived image
	Fingerprint Fingerprint `json:"fingerprint,omitempty"`

	// The number of layers by which this image differs from the associated image basis. @OutputOnly
	Distance int32 `json:"distance,omitempty"`

	// This contains layer-specific metadata, if populated it has length “distance” and is ordered with [distance] being the layer immediately following the base image and [1] being the final layer.
	LayerInfo []Layer `json:"layerInfo,omitempty"`

	// This contains the base image url for the derived image Occurrence @OutputOnly
	BaseResourceUrl string `json:"baseResourceUrl,omitempty"`
}
