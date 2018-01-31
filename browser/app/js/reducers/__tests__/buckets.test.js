/*
 * Minio Cloud Storage (C) 2018 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import reducer from "../buckets"
import * as actions from "../../actions/buckets"

describe("buckets reducer", () => {
  it("should return the initial state", () => {
    const initialState = reducer(undefined, {})
    expect(initialState).toEqual({
      list: [],
      filter: "",
      currentBucket: "",
      currentPrefix: ""
    })
  })

  it("should handle SET_BUCKETS", () => {
    const newState = reducer(undefined, {
      type: actions.SET_LIST,
      buckets: ["bk1", "bk2"]
    })
    expect(newState.list).toEqual(["bk1", "bk2"])
  })

  it("should handle SET_BUCKETS_FILTER", () => {
    const newState = reducer(undefined, {
      type: actions.SET_FILTER,
      filter: "test"
    })
    expect(newState.filter).toEqual("test")
  })

  it("should handle SELECT_BUCKET", () => {
    const newState = reducer(undefined, {
      type: actions.SET_CURRENT_BUCKET,
      bucket: "test"
    })
    expect(newState.currentBucket).toEqual("test")
  })

  it("should handle SET_CURRENT_PREFIX", () => {
    const newState = reducer(undefined, {
      type: actions.SET_CURRENT_PREFIX,
      prefix: "test"
    })
    expect(newState.currentPrefix).toEqual("test")
  })
})
