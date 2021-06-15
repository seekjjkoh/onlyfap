name := """play-scala-seed"""
organization := "com.example"

version := "0.0.1"

lazy val root = (project in file(".")).enablePlugins(PlayScala)

scalaVersion := "2.13.6"

libraryDependencies += guice
libraryDependencies += "org.scalatestplus.play" %% "scalatestplus-play" % "5.0.0" % Test

// Adds additional packages into Twirl
//TwirlKeys.templateImports += "com.example.controllers._"

// Adds additional packages into conf/routes
// play.sbt.routes.RoutesKeys.routesImport += "com.example.binders._"

enablePlugins(JavaAppPackaging)

javaOptions in Universal ++= Seq(
  "-Dpidfile.path=/dev/null"
)

dockerBaseImage := "openjdk:11"
