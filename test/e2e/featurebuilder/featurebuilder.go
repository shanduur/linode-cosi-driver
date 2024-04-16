package featurebuilder

import (
	"context"
	"strings"
	"testing"

	"sigs.k8s.io/e2e-framework/pkg/envconf"
	"sigs.k8s.io/e2e-framework/pkg/features"
)

type Label struct {
	Key, Value string
}

type LabelMatrix struct {
	Key         string
	ValueMatrix []string
}

type CommonAssesments map[string]func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context

func setupContextFromLabels(labels ...Label) func(ctx context.Context, t *testing.T, c *envconf.Config) context.Context {
	return func(ctx context.Context, t *testing.T, c *envconf.Config) context.Context {
		for _, l := range labels {
			ctx = context.WithValue(ctx, l.Key, l.Value)
		}

		return ctx
	}
}

func New(prefix string, labels ...Label) *features.FeatureBuilder {
	name := []string{}

	if prefix != "" {
		name = append(name, prefix)
	}

	for _, l := range labels {
		name = append(name, l.Value)
	}

	fb := features.New(strings.Join(name, "_"))

	for _, l := range labels {
		fb = fb.WithLabel(l.Key, l.Value)
	}

	return fb.WithSetup("fill the context", setupContextFromLabels(labels...))
}

func generateCombinations(
	labelIndex int,
	prefix string,
	labels []LabelMatrix,
	currentCombination []Label,
	combinations *[]*features.FeatureBuilder,
) {
	if labelIndex == len(labels) {
		name := []string{prefix}
		for _, l := range currentCombination {
			name = append(name, l.Value)
		}

		fb := features.New(strings.Join(name, "_"))
		for _, l := range currentCombination {
			fb = fb.WithLabel(l.Key, l.Value)
		}

		*combinations = append(*combinations, fb.WithSetup("fill the context", setupContextFromLabels(currentCombination...)))
		return
	}

	for _, value := range labels[labelIndex].ValueMatrix {
		currentCombination[labelIndex] = Label{Key: labels[labelIndex].Key, Value: value}
		generateCombinations(labelIndex+1, prefix, labels, currentCombination, combinations)
	}
}

func NewMatrix(
	prefix string,
	labels []LabelMatrix,
	assesments CommonAssesments,
) []*features.FeatureBuilder {
	var combinations []*features.FeatureBuilder

	generateCombinations(0, prefix, labels, make([]Label, len(labels)), &combinations)

	for _, fb := range combinations {
		for name, fn := range assesments {
			fb.Assess(name, fn)
		}
	}

	return combinations
}
