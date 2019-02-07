/*
Copyright 2018 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"log"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"sigs.k8s.io/controller-runtime/pkg/runtime/signals"

	"github.com/knative/eventing/contrib/kafka/pkg/dispatcher"
)

func main() {

	brokers := []string{"localhost:9094"}

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("unable to create logger: %v", err)
	}

	kafkaDispatcher, err := dispatcher.NewDispatcher(brokers, logger)
	if err != nil {
		logger.Fatal("unable to create kafka dispatcher.", zap.Error(err))
	}

	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()

	// Start both the manager (which notices ConfigMap changes) and the HTTP server.
	var g errgroup.Group

	g.Go(func() error {
		// Setups message receiver and blocks
		return kafkaDispatcher.Start(stopCh)
	})

	err = g.Wait()
	if err != nil {
		logger.Error("Either the kafka message receiver or the ConfigMap noticer failed.", zap.Error(err))
	}

}
