#
# spec file for package gettext-go-lint
#
# Copyright (c) 2024 SUSE LLC
#
# All modifications and additions to the file contributed by third parties
# remain the property of their copyright owners, unless otherwise agreed
# upon. The license for this file, and modifications and additions to the
# file, is the same license as for the pristine package itself (unless the
# license for the pristine package is not an Open Source License, in which
# case the license is the MIT License). An "Open Source License" is a
# license that conforms to the Open Source Definition (Version 1.9)
# published by the Open Source Initiative.

# Please submit bugfixes or comments via https://bugs.opensuse.org/
#


Name:           gettext-go-lint
Version:        0.0.1
Release:        0
Summary:        Check for common mistakes in gettext localizable strings in go code.
License:        Apache-2.0
# FIXME: use correct group, see "https://en.opensuse.org/openSUSE:Package_group_guidelines"
Group:          Development/Tools/Building
URL:            https://github.com/cbosdo/gettext-go-lint
Source0:        %{name}-%{version}.tar.gz
Source1:        vendor.tar.gz

%define go_bin  go

# Get the proper Go version on different distros
%if 0%{?suse_version}
BuildRequires:  golang(API) >= 1.20
%endif
# 0%{?suse_version}

%if 0%{?ubuntu}
%if 0%{?ubuntu} > 2204
%define go_version      1.22
%else
%define go_version      1.20
%endif
# 0%{?ubuntu} > 2204

BuildRequires:  golang-%{go_version}
%define go_bin  /usr/lib/go-%{go_version}/bin/go
%endif
# 0%{?ubuntu}

%if 0%{?debian}
BuildRequires:  golang >= 1.20
%endif
# 0%{?debian}


%if 0%{?fedora} || 0%{?rhel}
BuildRequires:  golang >= 1.19
%endif
# 0%{?fedora} || 0%{?rhel}

%description
A tool checking for gettext localized strings common mistakes in go source code.

%prep
%autosetup -p1 -a1

%build
%{go_bin} build -mod=vendor -buildmode=pie -ldflags="-X main.Version=%{version}"

%install
install -D -m0755 %{name} %{buildroot}/%{_bindir}/%{name}

%files
%license LICENSES/Apache-2.0.txt
%doc README.md
%{_bindir}/%{name}

%changelog

