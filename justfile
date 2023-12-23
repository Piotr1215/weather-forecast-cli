#!/usr/bin/env -S just --justfile
set export
set shell := ["bash", "-uc"]
                                    
# this list of available targets
# targets marked with * are main targets
default:
  @just --list --unsorted

prepare_release:
  git tag -a $(cat VERSION) -m "Release v$(cat VERSION)"
  goreleaser --snapshot --clean
  git tag

push_tag:
  git push origin $(cat VERSION)

release:
  goreleaser release --clean

