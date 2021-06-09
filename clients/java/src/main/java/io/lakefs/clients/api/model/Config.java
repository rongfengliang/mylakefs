/*
 * lakeFS API
 * lakeFS HTTP API
 *
 * The version of the OpenAPI document: 0.1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package io.lakefs.clients.api.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * Config
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class Config {
  public static final String SERIALIZED_NAME_BLOCKSTORE_NAMESPACE_EXAMPLE = "blockstore_namespace_example";
  @SerializedName(SERIALIZED_NAME_BLOCKSTORE_NAMESPACE_EXAMPLE)
  private String blockstoreNamespaceExample;

  public static final String SERIALIZED_NAME_BLOCKSTORE_NAMESPACE_VALIDITY_REGEX = "blockstore_namespace_ValidityRegex";
  @SerializedName(SERIALIZED_NAME_BLOCKSTORE_NAMESPACE_VALIDITY_REGEX)
  private String blockstoreNamespaceValidityRegex;

  public static final String SERIALIZED_NAME_WARNINGS = "warnings";
  @SerializedName(SERIALIZED_NAME_WARNINGS)
  private List<String> warnings = null;


  public Config blockstoreNamespaceExample(String blockstoreNamespaceExample) {
    
    this.blockstoreNamespaceExample = blockstoreNamespaceExample;
    return this;
  }

   /**
   * Get blockstoreNamespaceExample
   * @return blockstoreNamespaceExample
  **/
  @ApiModelProperty(required = true, value = "")

  public String getBlockstoreNamespaceExample() {
    return blockstoreNamespaceExample;
  }


  public void setBlockstoreNamespaceExample(String blockstoreNamespaceExample) {
    this.blockstoreNamespaceExample = blockstoreNamespaceExample;
  }


  public Config blockstoreNamespaceValidityRegex(String blockstoreNamespaceValidityRegex) {
    
    this.blockstoreNamespaceValidityRegex = blockstoreNamespaceValidityRegex;
    return this;
  }

   /**
   * Get blockstoreNamespaceValidityRegex
   * @return blockstoreNamespaceValidityRegex
  **/
  @ApiModelProperty(required = true, value = "")

  public String getBlockstoreNamespaceValidityRegex() {
    return blockstoreNamespaceValidityRegex;
  }


  public void setBlockstoreNamespaceValidityRegex(String blockstoreNamespaceValidityRegex) {
    this.blockstoreNamespaceValidityRegex = blockstoreNamespaceValidityRegex;
  }


  public Config warnings(List<String> warnings) {
    
    this.warnings = warnings;
    return this;
  }

  public Config addWarningsItem(String warningsItem) {
    if (this.warnings == null) {
      this.warnings = new ArrayList<String>();
    }
    this.warnings.add(warningsItem);
    return this;
  }

   /**
   * warnings to show user about this configuration
   * @return warnings
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "warnings to show user about this configuration")

  public List<String> getWarnings() {
    return warnings;
  }


  public void setWarnings(List<String> warnings) {
    this.warnings = warnings;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Config config = (Config) o;
    return Objects.equals(this.blockstoreNamespaceExample, config.blockstoreNamespaceExample) &&
        Objects.equals(this.blockstoreNamespaceValidityRegex, config.blockstoreNamespaceValidityRegex) &&
        Objects.equals(this.warnings, config.warnings);
  }

  @Override
  public int hashCode() {
    return Objects.hash(blockstoreNamespaceExample, blockstoreNamespaceValidityRegex, warnings);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Config {\n");
    sb.append("    blockstoreNamespaceExample: ").append(toIndentedString(blockstoreNamespaceExample)).append("\n");
    sb.append("    blockstoreNamespaceValidityRegex: ").append(toIndentedString(blockstoreNamespaceValidityRegex)).append("\n");
    sb.append("    warnings: ").append(toIndentedString(warnings)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

