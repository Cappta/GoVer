package gover

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestVersion(t *testing.T) {
	Convey("Given a base Version structure", t, func() {
		major := 1
		minor := 2
		build := 3
		revision := 4
		testBase := NewVersion(major, minor, build, revision)

		Convey("Then Major should be equal to provided Major", func() {
			So(testBase.GetMajor(), ShouldEqual, major)
		})
		Convey("Then Minor should be equal to provided Minor", func() {
			So(testBase.GetMinor(), ShouldEqual, minor)
		})
		Convey("Then Build should be equal to provided Build", func() {
			So(testBase.GetBuild(), ShouldEqual, build)
		})
		Convey("Then Revision should be equal to provided Revision", func() {
			So(testBase.GetRevision(), ShouldEqual, revision)
		})

		Convey("Given a version string", func() {
			baseString := fmt.Sprintf("%d.%d.%d.%d", major, minor, build, revision)
			Convey("When creating a version structure from string", func() {
				testBase, err := NewVersionFromString(baseString)

				Convey("Then error should be nil", func() {
					So(err, ShouldBeNil)
				})
				Convey("Then version structure should not be nil", func() {
					So(testBase, ShouldNotBeNil)
				})
				Convey("Then Major should be equal to provided Major", func() {
					So(testBase.GetMajor(), ShouldEqual, major)
				})
				Convey("Then Minor should be equal to provided Minor", func() {
					So(testBase.GetMinor(), ShouldEqual, minor)
				})
				Convey("Then Build should be equal to provided Build", func() {
					So(testBase.GetBuild(), ShouldEqual, build)
				})
				Convey("Then Revision should be equal to provided Revision", func() {
					So(testBase.GetRevision(), ShouldEqual, revision)
				})
			})
		})

		Convey("Given an invalid version string", func() {
			baseString := "1.2.3.d"
			Convey("When creating a version structure from string", func() {
				testBase, err := NewVersionFromString(baseString)

				Convey("Then error should not be nil", func() {
					So(err, ShouldNotBeNil)
				})
				Convey("Then version structure should be nil", func() {
					So(testBase, ShouldBeNil)
				})
			})
		})

		Convey("Given a Version structure with same values", func() {
			equalVersion := NewVersion(major, minor, build, revision)
			Convey("Then base should be equal", func() {
				So(testBase, VersionShouldEqual, equalVersion)
			})
		})
		Convey("Given a Version structure with a 5th zero value", func() {
			fiveValuesVersion := NewVersion(major, minor, build, revision, 0)
			Convey("Then base should be equal", func() {
				So(testBase, VersionShouldEqual, fiveValuesVersion)
			})
		})

		Convey("Given a Version structure with a higher Major", func() {
			higherMajor := NewVersion(major+1, minor, build, revision)
			Convey("Then base should be lower", func() {
				So(testBase, VersionShouldBeLowerThan, higherMajor)
			})
		})
		Convey("Given a Version structure with a higher Minor", func() {
			higherMinor := NewVersion(major, minor+1, build, revision)
			Convey("Then base should be lower", func() {
				So(testBase, VersionShouldBeLowerThan, higherMinor)
			})
		})
		Convey("Given a Version structure with a higher Build", func() {
			higherBuild := NewVersion(major, minor, build+1, revision)
			Convey("Then base should be lower", func() {
				So(testBase, VersionShouldBeLowerThan, higherBuild)
			})
		})
		Convey("Given a Version structure with a higher Revision", func() {
			higherRevision := NewVersion(major, minor, build, revision+1)
			Convey("Then base should be lower", func() {
				So(testBase, VersionShouldBeLowerThan, higherRevision)
			})
		})
		Convey("Given a Version structure with a 5th positive version value", func() {
			fiveValuesVersion := NewVersion(major, minor, build, revision, 1)
			Convey("Then base should be lower", func() {
				So(testBase, VersionShouldBeLowerThan, fiveValuesVersion)
			})
		})

		Convey("Given a Version structure with a lower Major", func() {
			lowerMajor := NewVersion(major-1, minor, build, revision)
			Convey("Then base should be greater", func() {
				So(testBase, VersionShouldBeGreaterThan, lowerMajor)
			})
		})
		Convey("Given a Version structure with a lower Minor", func() {
			lowerMinor := NewVersion(major, minor-1, build, revision)
			Convey("Then base should be greater", func() {
				So(testBase, VersionShouldBeGreaterThan, lowerMinor)
			})
		})
		Convey("Given a Version structure with a lower Build", func() {
			lowerBuild := NewVersion(major, minor, build-1, revision)
			Convey("Then base should be greater", func() {
				So(testBase, VersionShouldBeGreaterThan, lowerBuild)
			})
		})
		Convey("Given a Version structure with a lower Revision", func() {
			lowerRevision := NewVersion(major, minor, build, revision-1)
			Convey("Then base should be greater", func() {
				So(testBase, VersionShouldBeGreaterThan, lowerRevision)
			})
		})
		Convey("Given a Version structure with only 3 version values", func() {
			fifthValueVersion := NewVersion(major, minor, build)
			Convey("Then base should be greater", func() {
				So(testBase, VersionShouldBeGreaterThan, fifthValueVersion)
			})
		})
	})

}
