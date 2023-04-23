// Copyright 2023 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

// Package artifactcache provides a cache handler for act.
//
// Inspired by https://github.com/sp-ricard-valverde/github-act-cache-server
// and forked from: https://gitea.com/gitea/act_runner/src/branch/main/internal/app/artifactcache
//
// TODO: Authorization
// TODO: Restrictions for accessing a cache, see https://docs.github.com/en/actions/using-workflows/caching-dependencies-to-speed-up-workflows#restrictions-for-accessing-a-cache
// TODO: Force deleting cache entries, see https://docs.github.com/en/actions/using-workflows/caching-dependencies-to-speed-up-workflows#force-deleting-cache-entries

package artifactcache
