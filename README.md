Subnetting Tools
================

## Description

A simple set of tools for working with subnets
in golang.

## Commands

### calculateSubnets

> This command will calculate a list of subnet networks
> within the parent CIDR block.

**Syntax:** `calculateSubnet ${parentCIDR} ${subnetSize}`

* parentCidr = CIDR string (e.g. 10.11.0.0/16)
* subnetSize = size of the subnet (e.g. integer 0-32)
  but value must be within parent subnet.

