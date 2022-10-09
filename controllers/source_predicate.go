package controllers

import (
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	sourcev1 "github.com/fluxcd/source-controller/api/v1beta2"
)

type SourceRevisionChangePredicate struct {
	predicate.Funcs
}

func (SourceRevisionChangePredicate) Update(e event.UpdateEvent) bool {
	if e.ObjectOld == nil || e.ObjectNew == nil {
		return false
	}

	oldSource, ok := e.ObjectOld.(sourcev1.Source)
	if !ok {
		return false
	}

	newSource, ok := e.ObjectNew.(sourcev1.Source)
	if !ok {
		return false
	}

	if oldSource.GetArtifact() == nil && newSource.GetArtifact() != nil {
		return true
	}

	if oldSource.GetArtifact() != nil && newSource.GetArtifact() != nil &&
		oldSource.GetArtifact().Revision != newSource.GetArtifact().Revision {
		return true
	}

	return false
}

type SecretDeletePredicate struct {
}

// Create implements Predicate.
func (SecretDeletePredicate) Create(e event.CreateEvent) bool {
	return false
}

// Delete implements Predicate.
func (SecretDeletePredicate) Delete(e event.DeleteEvent) bool {
	return true
}

// Update implements Predicate.
func (SecretDeletePredicate) Update(e event.UpdateEvent) bool {
	if e.ObjectOld != nil && e.ObjectNew == nil {
		return true
	}

	return false
}

// Generic implements Predicate.
func (SecretDeletePredicate) Generic(e event.GenericEvent) bool {
	return false
}
