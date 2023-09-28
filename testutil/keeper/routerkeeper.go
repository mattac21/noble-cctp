/*
 * Copyright (c) 2023, © Circle Internet Financial, LTD.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package keeper

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MockRouterKeeper struct{}

func (MockRouterKeeper) HandleMessage(sdk.Context, []byte) error {
	return nil
}

type ErrorRouterKeeper struct{}

func (ErrorRouterKeeper) HandleMessage(sdk.Context, []byte) error {
	return errors.New("intentional error")
}
