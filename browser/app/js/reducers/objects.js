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

import * as actionsObjects from "../actions/objects"

export default (
  state = { list: [], sortBy: "", sortOrder: false, currentPrefix: "" },
  action
) => {
  switch (action.type) {
    case actionsObjects.SET_LIST:
      return {
        ...state,
        list: action.objects
      }
    case actionsObjects.SET_SORT_BY:
      return {
        ...state,
        sortBy: action.sortBy
      }
    case actionsObjects.SET_SORT_ORDER:
      return {
        ...state,
        sortOrder: action.sortOrder
      }
    case actionsObjects.SET_CURRENT_PREFIX:
      return {
        ...state,
        currentPrefix: action.prefix
      }
    default:
      return state
  }
}
